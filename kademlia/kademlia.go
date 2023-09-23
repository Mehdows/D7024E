package d7024e

import (
	"fmt"
	"net"
)



type Kademlia struct {
	me           Contact
	routingTable *RoutingTable
	network      *Network
	replicationFactor int
	k 		   		  int
	dictionary   map[string][]byte
}

func NewKademliaNode(address string) (kademlia Kademlia) {
	KademliaID := NewRandomKademliaID()
	kademlia.me = NewContact(KademliaID, address)
	kademlia.routingTable = NewRoutingTable(kademlia.me)
	kademlia.dictionary = make(map[string][]byte)
	kademlia.replicationFactor = 1
	kademlia.k = 3
	kademlia.network = &Network{&kademlia}
	go kademlia.network.Listen()
	return
}

func (Kademlia *Kademlia) JoinNetwork(address string) {
	contact := NewContact(NewRandomKademliaID(), address)
	Kademlia.network.SendPingMessage(&contact)
}

func (kademlia *Kademlia) LookupContact(target *Contact) (closestNode *Contact) {
	// list of k-closest nodes
	closestK := kademlia.routingTable.FindClosestContacts(target.ID, bucketSize)

	if len(closestK) == 0 {
		fmt.Println("No contacts found")
		return
	}

	closest := closestK[0]

	// find closest node
	for i := 0; i < len(closestK); i++ {
		if closestK[i].ID.CalcDistance(target.ID).Less(closest.ID.CalcDistance(target.ID)) {
			closest = closestK[i]
		}
	}

	// return closest node
	return &closest

}

func (kademlia *Kademlia) LookupData(hash string) {
	location := NewKademliaID(hash)
	recipient := kademlia.routingTable.FindClosestContacts(location, 1)
	go kademlia.network.SendFindDataMessage(recipient[0], hash)
}

func (kademlia *Kademlia) handleLookupData(message Message, conn net.Conn) {
	data := message.Data.(*findDataData)
	if kademlia.dictionary[data.Target.String()] != nil {
		kademlia.network.SendFindDataResponse(message, kademlia.dictionary[data.Target.String()], conn)
	} else {
		recipient := kademlia.routingTable.FindClosestContacts(data.Target, kademlia.k)
		kademlia.network.SendFindContactResponse(message, recipient, conn)
	}
}

func (kademlia *Kademlia) Store(data []byte) {
	location := NewKademliaID(string(data))
	recipient := kademlia.routingTable.FindClosestContacts(location, kademlia.replicationFactor)
	for i := 0; i < len(recipient); i++ {
		go kademlia.network.SendStoreMessage(recipient[i], location, data)
	}
}

func (kademlia *Kademlia) handleStore(message Message) {
	data := message.Data.(*storeData)
	kademlia.dictionary[data.Location.String()] = data.Data
}

func (Kademlia *Kademlia) Ping(id *KademliaID, address string) {
	Contact := NewContact(id, address)
	Kademlia.network.SendPingMessage(&Contact)
}

func (Kademlia *Kademlia) HandleRequest(conn net.Conn, message Message) {
	Kademlia.routingTable.AddContact(*message.sender)
	switch message.ID {
	case messageTypePing:
		Kademlia.network.SendPongMessage(message, conn)
	case messageTypeStore:
		Kademlia.handleStore(message)
	case messageTypeFindNode:
		// TODO
	case messageTypeFindValue:
		Kademlia.handleLookupData(message, conn)

	default:
		panic("Invalid request " + string(rune(message.ID)))
	}
}

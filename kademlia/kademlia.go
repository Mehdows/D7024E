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
	dictionary   map[string][]byte
}

func NewKademliaNode(address string) (kademlia Kademlia) {
	KademliaID := NewRandomKademliaID()
	kademlia.me = NewContact(KademliaID, address)
	kademlia.routingTable = NewRoutingTable(kademlia.me)
	kademlia.dictionary = make(map[string][]byte)
	kademlia.replicationFactor = 1
	kademlia.network = &Network{&kademlia}
	go kademlia.network.Listen()
	return
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
	
}


func (kademlia *Kademlia) Store(data []byte) {
	location := NewKademliaID(string(data))
	recipitent := kademlia.routingTable.FindClosestContacts(location, kademlia.replicationFactor)
	for i := 0; i < len(recipitent); i++ {
		go kademlia.network.SendStoreMessage(recipitent[i], location, data)
	}
}

func (kademlia *Kademlia) handleStore(message Message) {
	data := message.Data.(*storeDataData)
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
		// TODO
	case messageTypeFindNode:
		// TODO
	case messageTypeFindValue:

	default:
		panic("Invalid request " + string(rune(message.ID)))
	}
}

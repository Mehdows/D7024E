package kademlia

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net"
)

// Kademlia parameters
const alpha int = 1

type Kademlia struct {
	me                Contact
	routingTable      *RoutingTable
	network           *Network
	replicationFactor int
	k                 int
	dictionary        map[string][]byte
}

func NewKademliaNode(address string, me string) (kademlia Kademlia) {
	sha1 := sha1.Sum([]byte(me))
	key := hex.EncodeToString(sha1[:])
	id := NewKademliaID(key)
	kademlia.me = NewContact(id, address)
	kademlia.routingTable = NewRoutingTable(kademlia.me)
	kademlia.dictionary = make(map[string][]byte)
	kademlia.replicationFactor = 1
	kademlia.k = 1
	kademlia.network = &Network{&kademlia}
	go kademlia.network.Listen()
	return
}

func NewRandomKademliaNode(address string) (kademlia Kademlia) {
	id := NewRandomKademliaID()
	kademlia.me = NewContact(id, address)
	kademlia.routingTable = NewRoutingTable(kademlia.me)
	kademlia.dictionary = make(map[string][]byte)
	kademlia.replicationFactor = 1
	kademlia.k = 1
	kademlia.network = &Network{&kademlia}
	go kademlia.network.Listen()
	return
}

func (Kademlia *Kademlia) JoinNetwork(address string, id string) {
	sha1 := sha1.Sum([]byte(id))
	key := hex.EncodeToString(sha1[:])

	KademliaID := NewKademliaID(key)
	contact := NewContact(KademliaID, address)
	Kademlia.routingTable.AddContact(contact)
	Kademlia.LookupContact(Kademlia.me.ID)
}

func (kademlia *Kademlia) LookupContact(target *KademliaID) (closestNode *Contact) {
	net := kademlia.network

	// Create a shortlist for the search
	shortList := kademlia.routingTable.FindClosestContacts(target, alpha)

	closest := shortList[0]
	oldClose := shortList[0]

	for true {
		// Send alpha FIND_NODE RPCs
		response := net.SendFindContactMessage(&closest, target)

		shortList = append(shortList, response.Data.(*responseFindNodeData).Contacts...)
		closest = shortList[len(shortList)-1]

		fmt.Print("Cloest: ", closest.ID.String(), " OldClose: ", oldClose.ID.String(), "\n")
		if oldClose.ID.Equals(closest.ID) {
			fmt.Println("break")
			break
		} else {
			oldClose = closest
		}
	}
	return &closest
}

func (kademlia *Kademlia) handleLookUpContact(message Message, conn net.Conn) {
	data := message.Data.(*findNodeData)
	recipients := kademlia.routingTable.FindClosestContacts(&data.Target, kademlia.k)
	if len(recipients) == 0 {
		recipients = append(recipients, kademlia.me)
	}
	if kademlia.me.ID.Less(recipients[0].ID) {
		recipients[0] = kademlia.me
	}
	kademlia.network.SendFindContactResponse(message, recipients, conn)
}

func (kademlia *Kademlia) LookupData(hash string) {
	location := NewKademliaID(hash)
	recipient := kademlia.LookupContact(location)
	go kademlia.network.SendFindDataMessage(*recipient, hash)
}

func (kademlia *Kademlia) handleLookupData(message Message, conn net.Conn) {
	data := message.Data.(*findData)
	if kademlia.dictionary[data.Target.String()] != nil {
		kademlia.network.SendFindDataResponse(message, kademlia.dictionary[data.Target.String()], conn)
	} else {
		recipient := kademlia.routingTable.FindClosestContacts(&data.Target, kademlia.k)
		kademlia.network.SendFindContactResponse(message, recipient, conn)
	}
}

func (kademlia *Kademlia) Store(data []byte) {
	sha1 := sha1.Sum(data)
	key := hex.EncodeToString(sha1[:])

	location := NewKademliaID(key)
	recipient := kademlia.LookupContact(location)
	fmt.Println("Storing data at: ", location.String(), " on node: ", recipient.Address)
	go kademlia.network.SendStoreMessage(*recipient, location, []byte(data))
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
	defer Kademlia.routingTable.AddContact(message.Sender)
	switch message.ID {
	case messageTypePing:
		Kademlia.network.SendPongMessage(message, conn)
	case messageTypeStore:
		Kademlia.handleStore(message)
	case messageTypeFindNode:
		Kademlia.handleLookUpContact(message, conn)
	case messageTypeFindValue:
		Kademlia.handleLookupData(message, conn)

	default:
		panic("Invalid request " + string(rune(message.ID)))
	}
}

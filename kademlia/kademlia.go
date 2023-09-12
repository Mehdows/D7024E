package d7024e

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net"
)

type Kademlia struct {
	me           Contact
	routingTable *RoutingTable
	network      *Network
	dictionary   map[string][]byte
}

func NewKademliaNode(address string, ) (kademlia Kademlia) {
	KademliaID := NewRandomKademliaID()
	kademlia.me = NewContact(KademliaID, address)
	kademlia.routingTable = NewRoutingTable(kademlia.me)
	kademlia.dictionary = make(map[string][]byte)
	kademlia.network = &Network{&kademlia}
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
	// TODO
}

func (kademlia *Kademlia) Store(data []byte) {
	sha1 := sha1.Sum([]byte(data))
	key := hex.EncodeToString(sha1[:])
	fmt.Println("Stored data with key: ", key)
	fmt.Println("Stored hash: ", sha1)
}

func (Kademlia *Kademlia) Ping(id *KademliaID, address string) {
	Contact := NewContact(id, address)
	message := NewPingMessage(&Kademlia.me, &Contact)
	Kademlia.network.SendPingMessage(message)
}

func (Kademlia *Kademlia) HandleRequest(conn net.Conn, message Message) {
	switch message.ID {
	case messageTypePing:
		response := NewPongMessage(&Kademlia.me, message.sender)
		Kademlia.network.SendPongMessage(response, conn)
	case messageTypeStore:
		// TODO
	case messageTypeFindNode:
		// TODO
	case messageTypeFindValue:

	default:
		panic("Invalid request " + string(rune(message.ID)))
	}
}

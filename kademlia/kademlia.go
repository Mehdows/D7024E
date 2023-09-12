package d7024e

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type Kademlia struct {
	me 	 Contact
	routingTable *RoutingTable
	dictionary   map[string][]byte
}

func NewKademliaNode(adress string) (kademlia Kademlia) {
	KademliaID := NewKademliaID(adress)
	kademlia.me = NewContact(KademliaID, adress)
	kademlia.routingTable = NewRoutingTable(kademlia.me)
	kademlia.dictionary = make(map[string][]byte)
	return 
}

func (kademlia *Kademlia) LookupContact(target *Contact) (closestNode *Contact){
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

func HandleRequest(request *Network, function string) {
	switch function {
	case "ping":
		request.SendPongMessage()
	case "lookup_contact":
		// TODO
	case "lookup_data":
		// TODO
	case "store":
		// TODO
	default:
		fmt.Println(function)
		panic("Invalid request " + function)
	}
}

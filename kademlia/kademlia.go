package d7024e

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type Kademlia struct {
	routingTable *RoutingTable
	network      *Network
	dictionary   *map[string][]byte
}

func (kademlia *Kademlia) LookupContact(target *Contact) {
	// TODO
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

func (kademlia *Kademlia) Store(data []byte) {
	sha1 := sha1.Sum([]byte(data))
	key := hex.EncodeToString(sha1[:])
	kademlia.dict[key] = string(data)
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

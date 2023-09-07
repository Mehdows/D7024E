package d7024e

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type Kademlia struct {
	routingTable *RoutingTable
	network      *Network
	dict 	   map[string]string
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

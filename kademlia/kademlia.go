package d7024e

import "fmt"

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
	// TODO
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

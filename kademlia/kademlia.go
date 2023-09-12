package d7024e

import "net"

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

func (Kademlia *Kademlia) findKBucket() {
	// TODO
}

func (Kademlia *Kademlia) HandleRequest(conn net.Conn, message Message) {
	switch message.ID {
	case messageTypePing:
		Kademlia.network.SendPongMessage(conn)
	case messageTypeStore:
		// TODO
	case messageTypeFindNode:
		// TODO
	case messageTypeFindValue:

	default:
		panic("Invalid request " + string(message.ID))
	}
}

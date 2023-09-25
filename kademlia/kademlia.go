package d7024e

import (
	"net"
	"sync"
)

// Kademlia parameters
const alpha int = 1
var wg sync.WaitGroup

type Kademlia struct {
	me                Contact
	routingTable      *RoutingTable
	network           *Network
	replicationFactor int
	k                 int
	dictionary        map[string][]byte
}

func NewKademliaNode(address string) (kademlia Kademlia) {
	KademliaID := NewRandomKademliaID()
	kademlia.me = NewContact(KademliaID, address)
	kademlia.routingTable = NewRoutingTable(kademlia.me)
	kademlia.dictionary = make(map[string][]byte)
	kademlia.replicationFactor = 1
	kademlia.k = 1
	kademlia.network = &Network{&kademlia}
	go kademlia.network.Listen()
	return
}

func (Kademlia *Kademlia) JoinNetwork(address string, id byte) {
	KademliaID := KademliaID{id}
	contact := NewContact(&KademliaID, address)
	Kademlia.routingTable.AddContact(contact)
	Kademlia.LookupContact(&Kademlia.me)
}

func (kademlia *Kademlia) LookupContact(target *Contact) (closestNode *Contact) {
	net := kademlia.network
	previousClosestNode := kademlia.me
	
	// Create a shortlist for the search
	shortList := kademlia.routingTable.FindClosestContacts(target.ID, alpha)

	// Send alpha FIND_NODE RPCs
	response := net.SendFindContactMessage(shortList[0], target.ID)

	// Add the contacts from the response to the shortlist
	for i := 0; i < len(response.Data.(*responseFindNodeData).Contacts); i++ {
		shortList = append(shortList, response.Data.(*responseFindNodeData).Contacts[i])
	}

	// Find closest to target from shortlist
	for i := 0; i < len(shortList); i++ {
		closestNode := shortList[0]
		if shortList[i].ID.CalcDistance(target.ID).Less(target.ID.CalcDistance(closestNode.ID)) {
			closestNode = shortList[i]
		}
	}

	// If closest node is target, return closest node
	if closestNode == &previousClosestNode {
		return closestNode
	} else {
		// Else, repeat the process with the closest node
		previousClosestNode = *closestNode
		return kademlia.LookupContact(closestNode)
	}
}

func (kademlia *Kademlia) handleLookUpContact(message Message, conn net.Conn) {
	data := message.Data.(*findNodeData)
	recipient := kademlia.routingTable.FindClosestContacts(data.Target, kademlia.k)
	kademlia.network.SendFindContactResponse(message, recipient, conn)
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
		Kademlia.handleLookUpContact(message, conn)
	case messageTypeFindValue:
		Kademlia.handleLookupData(message, conn)

	default:
		panic("Invalid request " + string(rune(message.ID)))
	}
}

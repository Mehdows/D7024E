package d7024e

import (
	"net"
	"sync"
)

// Kademlia parameters
const alpha int = 3
var wg sync.WaitGroup

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
	net := kademlia.network
	// Create a channel for the responses
	resCh := make(chan []Contact, alpha)
	conCh := make(chan Contact, alpha)

	// Create a shortlist for the search
	shortList := kademlia.routingTable.FindClosestContacts(target.ID, bucketSize)

	// Send alpha FindContactMessages to alpha contacts in the shortlist
	if len(shortList) < alpha {
		for i := 0; i < len(shortList); i++ {
			wg.Add(1)
			go AsyncLookup(target.ID, shortList[i], *net, resCh, conCh)
		}
	}else {
		for i := 0; i < alpha; i++ {
			wg.Add(1)
			go AsyncLookup(target.ID, shortList[i], *net, resCh, conCh)
		}
	}

	// Wait for all the responses to arrive
	wg.Wait()
	close(resCh)
	close(conCh)

	// Create a list of all the responses
	var responses []Contact
	for response := range resCh {
		responses = append(responses, response...)
	}

	// Create a list of all the contacts
	var contacts []Contact
	for contact := range conCh {
		contacts = append(contacts, contact)
	}

	// Update the shortlist
	shortList = UpdateShortlist(shortList, responses, contacts[0])
	// return closest node
	return &closest
}

// AsyncLookup sends a FindContactMessage to the receiver and writes the response to a channel.
func AsyncLookup(targetID KademliaID, receiver Contact, net Network, ch chan []Contact, conCh chan Contact) {
	defer wg.Done()
	// Send the message and wait for the response
	response := net.SendFindContactMessage(targetID, receiver)

	// Write the response to the channel
	ch <- response
	conCh <- receiver
}

// UpdateShortlist updates the shortlist with the responses and the contact.
func (kademlia *Kademlia) UpdateShortlist (shortList []Contact, reslist []Contact, contact Contact) []Contact {
	// TODO
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

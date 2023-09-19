package d7024e

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
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

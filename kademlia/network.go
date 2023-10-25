package kademlia

import (
	"fmt"
	"net"
)

type Network struct {
	kademlia *Kademlia
}

func (network *Network) Listen() {
	address := network.kademlia.me.Address
	fmt.Println("Listening on: ", address)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go network.handleConnection(conn)
	}

}

func (network *Network) handleConnection(conn net.Conn) {
	data := make([]byte, 1024)
	_, err := conn.Read(data)

	if err != nil {
		panic(err)
	}
	var message Message
	DeserializeMessage(data, &message)

	network.kademlia.HandleRequest(conn, message)
}

// SendPingMessage sends a ping message to the contact
func (network *Network) SendPingMessage(reciever *Contact) string {
	me := network.kademlia.me
	message := NewPingMessage(me, *reciever)
	_ = network.dialAndSend(message)
	return "pong"

}

func (network *Network) SendPongMessage(pingMessage Message, conn net.Conn) {
	response := NewPongMessage(pingMessage)
	network.responseToConn(response, conn)
}

func (network *Network) SendFindContactMessage(receiver *Contact, hashToFind *KademliaID) Message {
	message := NewFindNodeMessage(network.kademlia.me, *receiver, *hashToFind)
	reply := network.dialAndSend(message)
	return reply
}

func (network *Network) SendFindContactResponse(message Message, contacts []Contact, conn net.Conn) {
	response := NewFindNodeResponse(network.kademlia.me, message.Sender, contacts)
	network.responseToConn(response, conn)
}

// SendFindDataMessage sends a find data message to the closest node to the hash
// returns the data if found, otherwise the closest contacts
func (network *Network) SendFindDataMessage(closestNode Contact, hash string) Message {
	message := NewFindValueMessage(network.kademlia.me, closestNode, *NewKademliaID(hash))

	response := network.dialAndSend(message)
	return response
}

// SendStoreMessage sends a store message to the closest node to the hash
func (network *Network) SendStoreMessage(receiver Contact, hash *KademliaID, data []byte) {
	datastruct := NewStoreData(*hash, data)
	message := NewStoreMessage(network.kademlia.me, receiver, datastruct)
	network.dialAndSend(message)
}

func (network *Network) SendFindDataResponse(message Message, data []byte, conn net.Conn) {
	response := NewFindValueResponse(network.kademlia.me, message.Sender, data)
	network.responseToConn(response, conn)
}

func (network *Network) dialAndSend(message Message) Message {
	data := SerializeMessage(&message)
	conn, err := net.Dial("tcp", message.Receiver.Address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write(data)

	//create byte buffer
	reply := network.listenForReply(conn)

	return reply
}

func (network *Network) listenForReply(conn net.Conn) Message {
	res := make([]byte, 1024)
	_, err := conn.Read(res)
	if err != nil {
		panic(err)
	}
	var message Message
	DeserializeMessage(res, &message)
	return message
}

func (network *Network) responseToConn(message Message, conn net.Conn) {
	data := SerializeMessage(&message)
	conn.Write(data)
}

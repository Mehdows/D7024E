package d7024e

import (
	"net"

)

type Network struct {
	kademlia *Kademlia
}

func (network *Network) Listen() {
	for {
		address := network.kademlia.me.Address
		ln, err := net.Listen("tcp", address)
		if err != nil {
			panic(err)
		}
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

	message := DeserializeMessage(data)

	network.kademlia.HandleRequest(conn, message)
}

// SendPingMessage sends a ping message to the contact
func (network *Network) SendPingMessage(reciever *Contact) string {
	me := network.kademlia.me
	message := NewPingMessage(&me, reciever)

	_ = network.dialAndSend(message)

	return "pong"

}

func (network *Network) SendPongMessage(pingMessage Message, conn net.Conn) {
	response := NewPongMessage(pingMessage)
	data := SerializeMessage(&response)
	conn.Write(data)
}

func (network *Network) SendFindContactMessage(receiver Contact, hashToFind *KademliaID) Message{
	message := NewFindNodeMessage(&network.kademlia.me, &receiver, hashToFind)
	reply := network.dialAndSend(message)
	return reply
}

// SendFindDataMessage sends a find data message to the closest node to the hash
// returns the data if found, otherwise the closest contacts
func (network *Network) SendFindDataMessage(closestNode Contact, hash string) Message{
	message := NewFindValueMessage(&network.kademlia.me, &closestNode, NewKademliaID(hash))
	response := network.dialAndSend(message)
	return response
}

// SendStoreMessage sends a store message to the closest node to the hash
func (network *Network) SendStoreMessage(receiver Contact, data []byte) {
	message := NewStoreMessage(&network.kademlia.me, &receiver, data)
	network.dialAndSend(message)
}

func (network *Network) dialAndSend(message Message) Message{
	data := SerializeMessage(&message)
	conn, err := net.Dial("tcp", message.receiver.Address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write(data)
	//create byte buffer
	reply := network.listenForReply(conn)
	
	return reply
}

func (network *Network) listenForReply(conn net.Conn) Message{
	res := make([]byte, 1024)

	_, err := conn.Read(res)
	if err != nil {
		panic(err)
	}
<<<<<<< HEAD
	return DeserializeMessage(res)
=======
	conn.Close()
	return "pong"

}

func (network *Network) SendPongMessage(message Message, conn net.Conn) {
	reciever := message.receiver
	message.receiver = message.sender
	message.sender = reciever
	message.ID = messageTypePing
	message.IsResponse = true

	data := SerializeMessage(&message)
	conn.Write(data)
}

func (network *Network) SendFindContactMessage(contact *Contact, target *KademliaID) []Contact {
	// TODO
	return nil
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
>>>>>>> 9b744b4530608d6934ce33b6406f64bdad23c259
}

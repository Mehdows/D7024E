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

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}

func (network *Network) dialAndSend(message Message) Message{
	data := SerializeMessage(&message)
	conn, err := net.Dial("tcp", message.receiver.Address)
	if err != nil {
		panic(err)
	}
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
	conn.Close()
	return DeserializeMessage(res)
}
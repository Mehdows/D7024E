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
func (network *Network) SendPingMessage(message Message) string {
	data := SerializeMessage(&message)
	conn, err := net.Dial("tcp", message.receiver.Address)
	if err != nil {
		panic(err)
	}
	conn.Write(data)
	//create byte buffer
	res := make([]byte, 1024)

	_, err = conn.Read(res)
	if err != nil {
		panic(err)
	}
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

func (network *Network) SendFindContactMessage(message Message) {
	

}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}

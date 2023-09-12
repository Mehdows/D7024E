package d7024e

import (
	"fmt"
	"net"
)

type Network struct {
	kademlia *Kademlia
}

func (network *Network) Listen(ip string, port int) {
	for {
		ln, err := net.Listen("tcp", ip+":"+fmt.Sprint(port))
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
func (network *Network) SendPingMessage(conn net.Conn) string {
	message := Message{
		ID:         messageTypePing,
		IsResponse: false,
	}
	data := SerializeMessage(&message)
	conn.Write(data)

	//create byte buffer
	res := make([]byte, 1024)

	_, err := conn.Read(res)
	if err != nil {
		panic(err)
	}
	conn.Close()
	return "pong"

}

func (network *Network) SendPongMessage(conn net.Conn) string {
	message := Message{
		ID:         messageTypePing,
		IsResponse: true,
	}
	data := SerializeMessage(&message)
	conn.Write(data)
}

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}

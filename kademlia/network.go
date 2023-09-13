package d7024e

import (
	"bytes"
	"fmt"
	"net"
	"strings"
)

type Network struct {
	connection *net.Conn
	contact    *Contact
	kademlia   *Kademlia
}

func Listen(ip string, port int) {
	for {
		ln, err := net.Listen("tcp", ip+":"+fmt.Sprint(port))
		if err != nil {
			panic(err)
		}
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	//read data from connection

	network := Network{}
	network.connection = &conn

	s := network.listenResponse()

	split := strings.Split(s, " ")

	hash := split[0]
	function := split[1]
	address := conn.RemoteAddr().String()
	kademliaID := NewKademliaID(hash)
	Contact := NewContact(kademliaID, address)
	network.contact = &Contact

	HandleRequest(&network, function)
}

func (network *Network) getContactConnection() {
	address := network.contact.Address
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	network.connection = &conn
}

func (network *Network) listenResponse() string {
	conn := *network.connection
	data := make([]byte, 1024)
	_, err := conn.Read(data)
	if err != nil {
		panic(err)
	}
	data = bytes.Trim(data, "\x00")
	dataString := string(data[:])
	return dataString
}

// SendPingMessage sends a ping message to the contact
func (network *Network) SendPingMessage() string {
	//Send ping message to contact with net
	network.getContactConnection()
	conn := *network.connection
	defer conn.Close()
	hash := network.contact.ID.String()
	conn.Write([]byte(hash + " ping"))

	dataString := network.listenResponse()
	pongMessage := strings.Split(dataString, " ")
	return pongMessage[1]
}

func (network *Network) SendPongMessage() {
	conn := *network.connection
	hash := network.contact.ID.String()
	conn.Write([]byte(hash + " pong"))
}

// SendFindContactMessage sends a find contact message to the
func (network *Network) SendFindContactMessage(contact *Contact) {
	
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}

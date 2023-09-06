package d7024e

import "net"

type Network struct {
}

func Listen(ip string, port int) {
	//TODO

}

// SendPingMessage sends a ping message to the contact
func (network *Network) SendPingMessage(contact *Contact) {
	address := contact.Address
	connClient, err := net.Dial("tcp", address)

	if err != nil {
		panic(err)
	}
	_, err = connClient.Write([]byte("ping"))
	defer connClient.Close()
}

// SendFindContactMessage sends a find contact message to the
func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}

package d7024e

import (
	"testing"
)

func TestSendPingMessage(t *testing.T) {

	Kademlia := NewKademliaNode("localhost:8080")
	go Kademlia.network.Listen()
	contact := NewContact(Kademlia.me.ID, "localhost:8080")
	res := Kademlia.network.SendPingMessage(&contact)

	if res != "pong" {
		t.Errorf("SendPingMessage() = %s; want pong", res)
	}
	
}

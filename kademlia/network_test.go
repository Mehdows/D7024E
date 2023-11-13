package kademlia

import "testing"

// not a unit test
func TestSendPingMessage(t *testing.T) {

	Kademlia := NewRandomKademliaNode("localhost:8080")
	contact := NewContact(Kademlia.me.ID, "localhost:8080")
	res := Kademlia.network.SendPingMessage(&contact)

	if res != "pong" {
		t.Errorf("SendPingMessage() = %s; want pong", res)
	}

}

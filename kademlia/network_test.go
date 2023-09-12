package d7024e

import (
	"testing"
)

func TestSendPingMessage(t *testing.T) {
	go Listen("localhost", 8080)
	net := Network{}
	contact := NewContact(NewRandomKademliaID(), "localhost:8080")
	net.contact = &contact
	res := net.SendPingMessage()
	if res != "pong" {
		t.Errorf("SendPingMessage() = %s; want pong", res)
	}
}

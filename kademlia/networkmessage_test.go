package d7024e

import (
	"testing"
)

func TestNewPingMessage(t *testing.T) {
	contact := NewContact(NewRandomKademliaID(), "localhost:8080")
	message := NewPingMessage(contact, contact)
	if message.Sender != contact {
		t.Errorf("NewPingMessage() = %s; want %s", message.Sender.String(), contact.String())
	}
	if message.Receiver.String() != contact.String() {
		t.Errorf("NewPingMessage() = %s; want %s", message.Receiver.String(), contact.String())
	}
	if message.ID != messageTypePing {
		t.Errorf("NewPingMessage() = %d; want %d", message.ID, messageTypePing)
	}
	if message.IsResponse != false {
		t.Errorf("NewPingMessage() = %t; want %t", message.IsResponse, false)
	}
}

func TestNewPongMessage(t *testing.T) {
	contact := NewContact(NewRandomKademliaID(), "localhost:8080")
	message := NewPingMessage(contact, contact)
	pong := NewPongMessage(message)
	if pong.Sender.String() != contact.String() {
		t.Errorf("NewPongMessage() = %s; want %s", pong.Sender.String(), contact.String())
	}
	if pong.Receiver.String() != contact.String() {
		t.Errorf("NewPongMessage() = %s; want %s", pong.Receiver.String(), contact.String())
	}
	if pong.ID != messageTypePing {
		t.Errorf("NewPongMessage() = %d; want %d", pong.ID, messageTypePing)
	}
	if pong.IsResponse != true {
		t.Errorf("NewPongMessage() = %t; want %t", pong.IsResponse, true)
	}
}

func TestNewFindNodeMessage(t *testing.T) {
	contact := NewContact(NewRandomKademliaID(), "localhost:8080")
	target := NewRandomKademliaID()
	message := NewFindNodeMessage(contact, contact, *target)
	if message.Sender.String() != contact.String() {
		t.Errorf("NewFindNodeMessage() = %s; want %s", message.Sender.String(), contact.String())
	}
	if message.Receiver != contact {
		t.Errorf("NewFindNodeMessage() = %s; want %s", message.Receiver.String(), contact.String())
	}
	if message.ID != messageTypeFindNode {
		t.Errorf("NewFindNodeMessage() = %d; want %d", message.ID, messageTypeFindNode)
	}
	if message.IsResponse != false {
		t.Errorf("NewFindNodeMessage() = %t; want %t", message.IsResponse, false)
	}
	if message.Data.(*findNodeData).Target != *target {
		t.Errorf("NewFindNodeMessage() = %s; want %s", message.Data.(*findNodeData).Target, target)
	}
}

func TestNewFindNodeResponse(t *testing.T) {
	contact := NewContact(NewRandomKademliaID(), "localhost:8080")
	recipient := NewContact(NewRandomKademliaID(), "localhost:8081")
	contacts := []Contact{contact}
	message := NewFindNodeResponse(contact, recipient, contacts)
	if message.Sender.String() != contact.String() {
		t.Errorf("NewFindNodeResponse() = %s; want %s", message.Sender.String(), contact.String())
	}
	if message.Receiver.String() != recipient.String() {
		t.Errorf("NewFindNodeResponse() = %s; want %s", message.Receiver.String(), contact.String())
	}
	if message.ID != messageTypeFindNode {
		t.Errorf("NewFindNodeResponse() = %d; want %d", message.ID, messageTypeFindNode)
	}
	if message.IsResponse != true {
		t.Errorf("NewFindNodeResponse() = %t; want %t", message.IsResponse, true)
	}
	if message.Data.(*responseFindNodeData).Contacts[0] != contact {
		t.Errorf("NewFindNodeResponse() = %v; want %v", message.Data.(*responseFindNodeData).Contacts[0], contact)
	}
}

func TestNewFindValueMessage(t *testing.T) {
	contact := NewContact(NewRandomKademliaID(), "localhost:8080")
	target := NewRandomKademliaID()
	message := NewFindValueMessage(contact, contact, *target)
	if message.Sender.String() != contact.String() {
		t.Errorf("NewFindValueMessage() = %s; want %s", message.Sender.String(), contact.String())
	}
	if message.Receiver.String() != contact.String() {
		t.Errorf("NewFindValueMessage() = %s; want %s", message.Receiver.String(), contact.String())
	}
	if message.ID != messageTypeFindValue {
		t.Errorf("NewFindValueMessage() = %d; want %d", message.ID, messageTypeFindValue)
	}
	if message.IsResponse != false {
		t.Errorf("NewFindValueMessage() = %t; want %t", message.IsResponse, false)
	}
	if message.Data.(*findData).Target != *target {
		t.Errorf("NewFindValueMessage() = %s; want %s", message.Data.(*findData).Target, target)
	}
}

func TestNewFindValueResponse(t *testing.T) {
	contact := NewContact(NewRandomKademliaID(), "localhost:8080")
	data := []byte("data")
	message := NewFindValueResponse(contact, contact, data)
	if message.Sender.String() != contact.String() {
		t.Errorf("NewFindValueResponse() = %s; want %s", message.Sender.String(), contact.String())
	}
	if message.Receiver.String() != contact.String() {
		t.Errorf("NewFindValueResponse() = %s; want %s", message.Receiver.String(), contact.String())
	}
	if message.ID != messageTypeFindValue {
		t.Errorf("NewFindValueResponse() = %d; want %d", message.ID, messageTypeFindValue)
	}
	if message.IsResponse != true {
		t.Errorf("NewFindValueResponse() = %t; want %t", message.IsResponse, true)
	}
	if string(message.Data.([]byte)) != string(data) {
		t.Errorf("NewFindValueResponse() = %s; want %s", string(message.Data.([]byte)), string(data))
	}
}

func TestNewStoreMessage(t *testing.T) {
	contact := NewContact(NewRandomKademliaID(), "localhost:8080")
	id := NewRandomKademliaID()
	data := NewStoreData(*id, []byte("data"))
	message := NewStoreMessage(contact, contact, data)
	if message.Sender != contact {
		t.Errorf("NewStoreMessage() = %s; want %s", message.Sender.String(), contact.String())
	}
	if message.Receiver.String() != contact.String() {
		t.Errorf("NewStoreMessage() = %s; want %s", message.Receiver.String(), contact.String())
	}
	if message.ID != messageTypeStore {
		t.Errorf("NewStoreMessage() = %d; want %d", message.ID, messageTypeStore)
	}
	if message.IsResponse != false {
		t.Errorf("NewStoreMessage() = %t; want %t", message.IsResponse, false)
	}
	if message.Data.(*storeData).Location != *id {
		t.Errorf("NewStoreMessage() = %s; want %s", message.Data.(*storeData).Location, id)
	}
	if string(message.Data.(*storeData).Data) != string(data.Data) {
		t.Errorf("NewStoreMessage() = %s; want %s", string(message.Data.(*storeData).Data), string(data.Data))
	}
}

func TestNewStoreResponse(t *testing.T) {
	contact := NewContact(NewRandomKademliaID(), "localhost:8080")
	message := newStoreResponseMessage(contact, contact, nil)
	if message.Sender.String() != contact.String() {
		t.Errorf("NewStoreResponse() = %s; want %s", message.Sender.String(), contact.String())
	}
	if message.Receiver.String() != contact.String() {
		t.Errorf("NewStoreResponse() = %s; want %s", message.Receiver.String(), contact.String())
	}
	if message.ID != messageTypeStore {
		t.Errorf("NewStoreResponse() = %d; want %d", message.ID, messageTypeStore)
	}
	if message.IsResponse != true {
		t.Errorf("NewStoreResponse() = %t; want %t", message.IsResponse, true)
	}
	if message.Error != nil {
		t.Errorf("NewStoreResponse() = %v; want %v", message.Error, nil)
	}
}

func TestNewStoreData(t *testing.T) {
	id := NewRandomKademliaID()
	data := NewStoreData(*id, []byte("data"))
	if data.Location != *id {
		t.Errorf("NewStoreData() = %s; want %s", data.Location, id)
	}
	if string(data.Data) != "data" {
		t.Errorf("NewStoreData() = %s; want %s", string(data.Data), "data")
	}
	if data.DataLength != len([]byte("data")) {
		t.Errorf("NewStoreData() = %d; want %d", data.DataLength, len([]byte("data")))
	}
}

func TestSerializeMessage(t *testing.T) {

	var deserialized Message
	data := NewStoreData(*NewRandomKademliaID(), []byte("data"))
	contact := NewContact(NewRandomKademliaID(), "localhost:8080")
	message := NewStoreMessage(contact, contact, data)
	serialized := SerializeMessage(&message)
	DeserializeMessage(serialized, &deserialized)
	var stData storeData
	stData.FillStruct(deserialized.Data.(map[string]interface{}))
	deserialized.Data = stData

	if deserialized.Sender.String() != message.Sender.String() {
		t.Errorf("SerializeMessage() = %s; want %s", deserialized.Sender.String(), message.Sender.String())
	}
	if deserialized.Receiver.String() != message.Receiver.String() {
		t.Errorf("SerializeMessage() = %s; want %s", deserialized.Receiver.String(), message.Receiver.String())
	}
	if deserialized.ID != message.ID {
		t.Errorf("SerializeMessage() = %d; want %d", deserialized.ID, message.ID)
	}
	if deserialized.IsResponse != message.IsResponse {
		t.Errorf("SerializeMessage() = %t; want %t", deserialized.IsResponse, message.IsResponse)
	}
	if deserialized.Data.(storeData).Location != message.Data.(storeData).Location {
		t.Errorf("SerializeMessage() = %s; want %s", deserialized.Data.(storeData).Location, message.Data.(storeData).Location)
	}
	if string(deserialized.Data.(storeData).Data) != string(message.Data.(storeData).Data) {
		t.Errorf("SerializeMessage() = %s; want %s", string(deserialized.Data.(storeData).Data), string(message.Data.(storeData).Data))
	}
	if deserialized.Data.(storeData).DataLength != message.Data.(storeData).DataLength {
		t.Errorf("SerializeMessage() = %d; want %d", deserialized.Data.(storeData).DataLength, message.Data.(storeData).DataLength)
	}
}

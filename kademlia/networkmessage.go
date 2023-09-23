package d7024e

import (
	"bytes"
	"encoding/json"
)

const (
	messageTypePing      = iota
	messageTypeStore     = iota
	messageTypeFindNode  = iota
	messageTypeFindValue = iota
)

type Message struct {
	sender     *Contact
	receiver   *Contact
	ID         int
	IsResponse bool
	Error 	   error
	Data       interface{}
}

type findNodeData struct {
	Target *KademliaID
}

type findDataData struct {
	Target *KademliaID
}

type storeData struct {
	Location *KademliaID
	Data []byte
	DataLength int
}

type responseFindNodeData struct {
	Contacts []Contact
}

func NewPingMessage(sender *Contact, receiver *Contact) Message {
	return Message{
		sender:     sender,
		receiver:   receiver,
		ID:         messageTypePing,
		IsResponse: false,
	}
}

func NewPongMessage(pingMessage Message) Message {
	return Message{
		sender:     pingMessage.receiver,
		receiver:   pingMessage.sender,
		ID:         messageTypePing,
		IsResponse: true,
	}
}

func NewFindNodeMessage(sender *Contact, receiver *Contact, target *KademliaID) Message {
	return Message{
		sender:     sender,
		receiver:   receiver,
		ID:         messageTypeFindNode,
		IsResponse: false,
		Data:       &findNodeData{target},
	}
}

func NewFindNodeResponse(sender *Contact, receiver *Contact, contacts []Contact) Message {
	return Message{
		sender:     sender,
		receiver:   receiver,
		ID:         messageTypeFindNode,
		IsResponse: true,
		Data:       &responseFindNodeData{contacts},
	}
}

func NewFindValueMessage(sender *Contact, receiver *Contact, target *KademliaID) Message {
	return Message{
		sender:     sender,
		receiver:   receiver,
		ID:         messageTypeFindValue,
		IsResponse: false,
		Data:       &findDataData{target},
	}
}

func NewFindValueResponse(sender *Contact, receiver *Contact, data []byte) Message{
	return Message{
		sender:     sender,
		receiver:   receiver,
		ID:         messageTypeFindValue,
		IsResponse: true,
		Data:       data,
	}
}

func NewStoreMessage(sender *Contact, receiver *Contact, data *storeData) Message {
	return Message{
		sender:     sender,
		receiver:   receiver,
		ID:         messageTypeStore,
		IsResponse: false,
		Data:       data,
	}
}

func newStoreResponseMessage(sender *Contact, receiver *Contact, err error) Message{
	return Message{
		sender:     sender,
		receiver:   receiver,
		ID:         messageTypeStore,
		IsResponse: true,
		Error:      err,
	}
}

func NewStoreData(location *KademliaID, data []byte) storeData {
	return storeData{
		Location: location,
		Data:     data,
		DataLength: len(data),
	}
}


// implement serialization with marshal
func SerializeMessage(message *Message) []byte {

	data, err := json.Marshal(message)

	if err != nil {
		panic(err)
	}

	return data
}

func DeserializeMessage(data []byte) Message {
	//remove empty bytes
	data = bytes.Trim(data, "\x00")
	err := json.Unmarshal(data, &Message{})
	if err != nil {
		panic(err)
	}

	return Message{}
}

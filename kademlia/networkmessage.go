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
	Data       interface{}
}

type findNodeData struct {
	Target *KademliaID
}

type findDataData struct {
	Target *KademliaID
}

type storeDataData struct {
	Data []byte
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

func NewFindContactMessage(sender *Contact, receiver *Contact, target *KademliaID) Message {
	return Message{
		sender:     sender,
		receiver:   receiver,
		ID:         messageTypeFindNode,
		IsResponse: false,
		Data:       &findNodeData{target},
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

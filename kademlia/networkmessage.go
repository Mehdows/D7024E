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

func NewFindValueMessage(sender *Contact, receiver *Contact, target *KademliaID) Message {
	return Message{
		sender:     sender,
		receiver:   receiver,
		ID:         messageTypeFindValue,
		IsResponse: false,
		Data:       &findDataData{target},
	}
}

func NewStoreMessage(sender *Contact, receiver *Contact, data *storeDataData) Message {
	return Message{
		sender:     sender,
		receiver:   receiver,
		ID:         messageTypeStore,
		IsResponse: false,
		Data:       data,
	}
}

func NewStoreData(location *KademliaID, data []byte) storeDataData {
	return storeDataData{
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

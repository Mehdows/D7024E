package kademlia

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
	Sender     Contact     `json:"Sender"`
	Receiver   Contact     `json:"Receiver"`
	ID         int         `json:"ID"`
	IsResponse bool        `json:"IsResponse"`
	Error      error       `json:"Error"`
	Data       interface{} `json:"Data"`
}

type findNodeData struct {
	Target KademliaID `json:"Target"`
}

type findData struct {
	Target KademliaID `json:"Target"`
}

type storeData struct {
	Location   KademliaID `json:"Location"`
	Data       []byte     `json:"Data"`
	DataLength int        `json:"DataLength"`
}

type responseFindNodeData struct {
	Contacts []Contact `json:"Contacts"`
}

func NewPingMessage(Sender Contact, Receiver Contact) Message {
	return Message{
		Sender:     Sender,
		Receiver:   Receiver,
		ID:         messageTypePing,
		IsResponse: false,
	}
}

func NewPongMessage(pingMessage Message) Message {
	return Message{
		Sender:     pingMessage.Receiver,
		Receiver:   pingMessage.Sender,
		ID:         messageTypePing,
		IsResponse: true,
	}
}

func NewFindNodeMessage(Sender Contact, Receiver Contact, target KademliaID) Message {
	return Message{
		Sender:     Sender,
		Receiver:   Receiver,
		ID:         messageTypeFindNode,
		IsResponse: false,
		Data:       &findNodeData{target},
	}
}

func NewFindNodeResponse(Sender Contact, Receiver Contact, contacts []Contact) Message {
	return Message{
		Sender:     Sender,
		Receiver:   Receiver,
		ID:         messageTypeFindNode,
		IsResponse: true,
		Data:       &responseFindNodeData{contacts},
	}
}

func NewFindValueMessage(Sender Contact, Receiver Contact, target KademliaID) Message {
	return Message{
		Sender:     Sender,
		Receiver:   Receiver,
		ID:         messageTypeFindValue,
		IsResponse: false,
		Data:       findData{target},
	}
}

func NewFindValueResponse(Sender Contact, Receiver Contact, data []byte) Message {
	return Message{
		Sender:     Sender,
		Receiver:   Receiver,
		ID:         messageTypeFindValue,
		IsResponse: true,
		Data:       data,
	}
}

func NewStoreMessage(Sender Contact, Receiver Contact, data storeData) Message {
	return Message{
		Sender:     Sender,
		Receiver:   Receiver,
		ID:         messageTypeStore,
		IsResponse: false,
		Data:       data,
	}
}

func newStoreResponseMessage(Sender Contact, Receiver Contact, err error) Message {
	return Message{
		Sender:     Sender,
		Receiver:   Receiver,
		ID:         messageTypeStore,
		IsResponse: true,
		Error:      err,
	}
}

func NewStoreData(location KademliaID, data []byte) storeData {
	return storeData{
		Location:   location,
		Data:       data,
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

func DeserializeMessage(data []byte, message *Message) {
	//remove empty bytes
	data = bytes.Trim(data, "\x00")
	err := json.Unmarshal(data, message)
	if err != nil {
		panic(err)
	}
	deserializeDataField(data, message)
}

func deserializeDataField(data []byte, message *Message) {
	var datastruct interface{}
	switch message.ID {
	case messageTypeFindNode:
		datastruct = new(findNodeData)
	case messageTypeFindValue:
		if message.IsResponse {
			datastruct = new(responseFindNodeData)
		} else {
			datastruct = new(findData)
		}
	case messageTypeStore:
		datastruct = new(storeData)
	}
	message.Data = datastruct
	err := json.Unmarshal(data, message)
	if err != nil {
		panic(err)
	}
}

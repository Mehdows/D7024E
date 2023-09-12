package d7024e

import (
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

// implement serialization with marshal
func SerializeMessage(message *Message) []byte {

	data, err := json.Marshal(message)

	if err != nil {
		panic(err)
	}

	return data
}

func DeserializeMessage(data []byte) Message {
	//make buffer
	err := json.Unmarshal(data, &Message{})
	if err != nil {
		panic(err)
	}

	return Message{}
}

package d7024e

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

const (
	messageTypePing      = iota
	messageTypeStore     = iota
	messageTypeFindNode  = iota
	messageTypeFindValue = iota
)

type Message struct {
	Sender     Contact     `json:"sender"`
	Receiver   Contact     `json:"receiver"`
	ID         int         `json:"id"`
	IsResponse bool        `json:"isResponse"`
	Error      error       `json:"error"`
	Data       interface{} `json:"data"`
}

type findNodeData struct {
	Target KademliaID `json:"target"`
}

type findData struct {
	Target KademliaID `json:"target"`
}

type storeData struct {
	Location   KademliaID `json:"location"`
	Data       []byte     `json:"data"`
	DataLength int        `json:"dataLength"`
}

type responseFindNodeData struct {
	Contacts []Contact `json:"contacts"`
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
}

func (s *storeData) FillStruct(m map[string]interface{}) error {
	for k, v := range m {
		err := SetField(s, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}

	structFieldValue.Set(val)
	return nil
}

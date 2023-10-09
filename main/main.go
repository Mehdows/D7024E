package main

import (
	"fmt"
	"github.com/Mehdows/D7024E/kademlia"
)

func main () {
	fmt.Println("Hello world!")
	Kademlia := NewKademliaNode("localhost:8080")
	Cli_handler(&Kademlia)
}
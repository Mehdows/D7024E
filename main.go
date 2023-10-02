package main

import (
	
)

func main () {
	Kademlia := NewKademliaNode("localhost:8080")
	Cli_handler(&Kademlia)
}
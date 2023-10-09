package main

import (
	"github.com/Mehdows/D7024E/kademlia"
)

func main() {

	Kademlia := kademlia.NewKademliaNode("localhost:80")
	kademlia.Cli_handler(&Kademlia)
}

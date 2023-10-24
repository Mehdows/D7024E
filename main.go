package main

import (
	"encoding/hex"
	"os"

	"github.com/Mehdows/D7024E/kademlia"
)

func main() {

	Kademlia := kademlia.NewKademliaNode("localhost:8080")

	if os.Args[1] == "CLI" {
		kademlia.Cli_handler(&Kademlia)
	}
	if len(os.Args) >= 2 {
		//convert string to byte arr
		id, _ := hex.DecodeString(os.Args[1])
		Kademlia.JoinNetwork(os.Args[2], id[0])
	}
}

package kademlia

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func Cli_handler(kademlia *Kademlia) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(kademlia.me.ID.String() + "> ")
	fmt.Println(kademlia.me.Address + "> ")
	for {
		//print bucket content
		for i := 0; i < IDLength*8; i++ {
			if kademlia.routingTable.buckets[i].list.Len() == 0 {
				continue
			}
			fmt.Println("Bucket ", i, ": ")
			for i := kademlia.routingTable.buckets[i].list.Front(); i != nil; i = i.Next() {
				fmt.Println("Contact: ", i.Value)
			}
		}

		fmt.Print("Enter Text: ")
		// reads user input until \n by default
		scanner.Scan()
		// handle error
		if scanner.Err() != nil {
			fmt.Println("Error: ", scanner.Err())
		}
		// Holds the string that was scanned
		text := scanner.Text()
		// Holds the first word of the string
		res := strings.Split(text, " ")
		choice := res[0]
		// If the user wants to get a file write its hash
		if choice == "get" && len(res) == 2 && len(res[1]) == 40 {
			fmt.Println("I will get a file with hash: ", res[1])
			res := kademlia.LookupData(res[1])
			if res == "" {
				fmt.Println("File not found")
			} else {
				fmt.Println("File content: ", res)
			}
			// If the user wants to store a file write its value
		} else if choice == "put" {
			newres := strings.Join(res[1:], " ")
			sha1 := sha1.Sum([]byte(newres))
			res := hex.EncodeToString(sha1[:])
			id := NewKademliaID(res)
			kademlia.Store([]byte(newres))
			fmt.Println("I will store a file with value: ", newres, " with hash: ", id.String())
		} else if choice == "stored" {
			for key, value := range kademlia.dictionary {
				fmt.Println("Key: ", key, " Value: ", value)
			}
		} else if choice == "exit" {
			fmt.Println("Exiting...")
			break

		} else {
			fmt.Println("Unknown command")
		}
	}

}

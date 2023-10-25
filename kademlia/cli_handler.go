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
		fmt.Println("You entered: ", text)
		// Holds the first word of the string
		res := strings.Split(text, " ")
		choice := res[0]
		fmt.Println("You entered: ", choice)
		// If the user wants to get a file write its hash
		if choice == "get" && len(res) == 2 {
			fmt.Println("I will get a file with hash: ", res[1])
			kademlia.LookupData(res[1])

			// If the user wants to store a file write its value
		} else if choice == "put" {
<<<<<<< HEAD
=======
			fmt.Println("1")
>>>>>>> daa4ba3450ab58846a9f4869f1b9909e2777b820
			newres := strings.Join(res[1:], " ")
			fmt.Println("2")
			sha1 := sha1.Sum([]byte(newres))
			fmt.Println("3")
			newres = hex.EncodeToString(sha1[:])
			fmt.Println("4")
			id := NewKademliaID(newres)
			fmt.Println("5")
			kademlia.Store([]byte(newres))
			fmt.Println("6")
			fmt.Println("I will store a file with value: ", res[1:], " with hash: ", id.String())
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

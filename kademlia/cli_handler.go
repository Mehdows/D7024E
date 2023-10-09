package kademlia

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func Cli_handler(kademlia *Kademlia) {
	scanner := bufio.NewScanner(os.Stdin)
	print(kademlia.me.Address + "> ")
	for {
		fmt.Print("Enter Text: ")
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		// Holds the first word of the string
		res := strings.Split(text, " ")
		choice := res[0]

		// If the user wants to get a file write its hash
		if choice == "get" && len(res) == 2 {
			fmt.Println("I will get a file with hash: ", res[1])
			kademlia.LookupData(res[1])

		// If the user wants to store a file write its value
		} else if choice == "put" {
			fmt.Println("I will store a file with value: ", res[1:])
			newres := strings.Join(res[1:], " ")
			bytearr := []byte(newres)
			kademlia.Store(bytearr)

		} else if choice == "exit" {
			fmt.Println("Exiting...")
			break

		} else {
			fmt.Println("Unknown command")
		}
	}

	// handle error
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}
}

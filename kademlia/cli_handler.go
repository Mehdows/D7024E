package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main () {
	Cli_handler()
}

func Cli_handler() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter Text: ")
		// reads user input until \n by default
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()
		// Holds the first word of the string
		res := strings.Split(text, " ")
		choice := res[0]

		if choice == "get" && len(res) == 2 {
			fmt.Println("I will get a file with hash: ", res[1])

		} else if choice == "put" {
			fmt.Println("I will store a file with value: ", res[1:])
			newres := strings.Join(res[1:], " ")
			bytearr := []byte(newres)
			fmt.Println(bytearr)
			Store(bytearr)

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

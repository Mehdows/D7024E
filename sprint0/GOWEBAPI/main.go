package main

import (
	
)


var SERVER_PORT string 
var CLIENT_ADDRESS string // possibly = "http://localhost:8081"

var packets_recieved int
var packets_sent int

type Response struct {
    someString string
}

var response Response;

var recieved int
var sent int


func setResponse() {
	response = Response{"Success!"}
}

func sendPacket(port int) {
	packets_sent++
	c := newClient(CLIENT_ADDRESS)
	setResponse()

	res := exchangeJson(c, response)
	
}


func main() {
	
}
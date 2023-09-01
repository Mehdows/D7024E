package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func listen(rw http.ResponseWriter, req *http.Request) {
	Body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(Body)

	respond(rw, response)
}

func respond(rw http.ResponseWriter, responseStruct interface{}) {

	fmt.Println("Sending response...")

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(responseStruct)
}

func SetupServer() {
	port := SERVER_PORT
	funcHandle := "/"
	http.HandleFunc(funcHandle, listen)
	http.ListenAndServe(port, nil)
}

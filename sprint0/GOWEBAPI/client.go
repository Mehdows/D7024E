package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"io"
	"net/http"
	//forms "arrowhead/Orchestrator/forms"
)

//Client struct
type client struct {
	httpAdrs string
}

//Init Client
func newClient(httpAdrs string) *client {
	c := client{httpAdrs: httpAdrs}
	return &c
}

func exchangeJson(c *client, struc interface{}) []byte {
	jsonStr, err := json.Marshal(struc)
	errorHandler(err)
	resp, err := http.Post(c.httpAdrs, "application/json", bytes.NewBuffer(jsonStr))
	errorHandler2(err)
	defer resp.Body.Close()

	jsonBody, err := io.ReadAll(resp.Body)

	fmt.Println(string(jsonBody))

	errorHandler(err)

	return jsonBody
}

func errorHandler(err error) {
	if err != nil {
		panic(err)
	}
}

func errorHandler2(err error) {
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println("No ServiceRegistry Found")
		fmt.Println("")
	}
}
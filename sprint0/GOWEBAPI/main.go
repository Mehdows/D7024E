package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Response struct {
    someString string
}

var recieved int
var sent int



func sendPacket(port int) {
	struc := &Response{"Sent Success"}
	jsonStr, err := json.Marshal(struc)
	if err != nil {
		fmt.Println("Error")
	}
	resp, err := http.Post(fmt.Sprint("localhost:", port), "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println("Error")
	}
	defer resp.Body.Close()
	strjson, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error")
	}
	res := ""
	json.Unmarshal(strjson, &res )


}
func listen(rw http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(body, &recieved)


	respond(rw, "Recieved Success")
}

func respond(rw http.ResponseWriter, str string) {

	fmt.Println("Sending response...")

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(str)
}

func SetupServer(port string) {
	funcHandle := "/sendPacket"
	http.HandleFunc(funcHandle, listen)
	http.ListenAndServe(port, nil)
}

func main() {
	
}
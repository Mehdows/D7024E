package main

import (
	"html/template"
	"net/http"
)

var SERVER_PORT string
var CLIENT_ADDRESS string // possibly = "http://localhost:8081"

var packets_recieved int
var packets_sent int

type Response struct {
	someString string
}

var response Response

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
	print(res)

}

func ping(w http.ResponseWriter, r *http.Request) {
	var filename = "index.html"
	t, err := template.ParseFiles(filename)
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, filename, nil)
}

func main() {
	http.HandleFunc("/", ping)
	// http handle function for index.css
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	http.ListenAndServe(":8080", nil)

}

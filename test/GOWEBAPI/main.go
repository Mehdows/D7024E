package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	request1()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Me!")
	fmt.Println("Endpoint Hit: aboutMe")
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutMe", aboutMe)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

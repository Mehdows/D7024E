package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	setupRoutes()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Me!")
	fmt.Println("Endpoint Hit: aboutMe")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the IP address from the query parameter
	ip := r.URL.Query().Get("ip")

	if ip == "" {
		http.Error(w, "IP address is required", http.StatusBadRequest)
		return
	}

	// Create a new ping instance
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		http.Error(w, "Error creating ping instance", http.StatusInternalServerError)
		return
	}

	// Set the ping options
	pinger.Count = 4
	pinger.Timeout = 5 * time.Second

	// Run the ping
	err = pinger.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Ping error: %v", err), http.StatusInternalServerError)
		return
	}

	// Get and send the ping statistics as the response
	stats := pinger.Statistics()
	fmt.Fprintf(w, "Ping Results for %s:\n", ip)
	fmt.Fprintf(w, "Packets Sent: %d\n", stats.PacketsSent)
	fmt.Fprintf(w, "Packets Received: %d\n", stats.PacketsRecv)
	fmt.Fprintf(w, "Packet Loss: %v%%\n", stats.PacketLoss)
	fmt.Fprintf(w, "Round-trip Min/Avg/Max = %v/%v/%v\n", stats.MinRtt, stats.AvgRtt, stats.MaxRtt)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutMe", aboutMe)
	http.HandleFunc("/ping", pingHandler) // Endpoint for handling ping requests
	log.Fatal(http.ListenAndServe(":8080", nil))
}
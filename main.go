package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/Mehdows/D7024E/kademlia"
)

const IP_PREFIX = ""
const BOOTSTRAP_IP = "172.18.0.2"
const BOOTSTRAP_ID = "0000000000000000000000000000000000000000"

func main() {
	ip, _ := getMyIP()
	Kademlia := kademlia.NewKademliaNode(ip)
	if !isBootstrap() {
		Kademlia.JoinNetwork(BOOTSTRAP_IP, BOOTSTRAP_ID)
	}

	kademlia.Cli_handler(&Kademlia)

}

func isBootstrap() bool {
	ip, err := getMyIP()
	if err != nil {
		log.Fatal(err)
	}
	return ip == BOOTSTRAP_IP
}

func getMyIP() (string, error) {
	containerName, _ := os.Hostname()
	ips, err := net.LookupIP(containerName)
	if err != nil {
		fmt.Println("Unknown host")
	} else {
		for _, ip := range ips {
			if strings.HasPrefix(ip.String(), IP_PREFIX) {
				return ip.String(), nil
			}
		}
	}
	return "", errors.New("Missing IP")
}

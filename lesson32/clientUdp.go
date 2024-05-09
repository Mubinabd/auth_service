package main

import (
	"fmt"
	"net"
)

func main() {
	sAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, sAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("Helloooo, im from good"))
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024)
	_, _, err = conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("error while readFromUdp", err)
		return
	}

	fmt.Println("message from server:", string(buffer))
}

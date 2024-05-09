package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Listening udp")

	buffer := make([]byte, 1024)

	for {
		_, adr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("error while readFromUdp", err)
			continue
		}

		fmt.Printf("addres from %s, message:%s", adr.String(), string(buffer))

		_, err = conn.WriteToUDP([]byte("udp receive your message"), adr)
		if err != nil {
			fmt.Println("error while writeToUdp", err)
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving UDP address: %s\n", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error dialing UDP: %s\n", err)
		return
	}
	defer conn.Close()

	go listenForMessages(conn)
	fmt.Println("Send your messages here: ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to send message: %s\n", err)
			break
		}
	}
}

func listenForMessages(conn *net.UDPConn) {
	for {
		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading UDP message: %s\n", err)
			return
		}
		fmt.Printf("%v\n", string(buf[:n]))
	}
}

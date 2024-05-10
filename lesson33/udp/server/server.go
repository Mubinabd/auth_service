package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

var (
	clientAddrs = make(map[string]*net.UDPAddr)
	mutex       sync.Mutex
)

func main() {
	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to set up UDP server: %s\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("UDP server listening on port 8080")

	buf := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read from UDP: %s\n", err)
			continue
		}
		message := remoteAddr.String() + ": " + string(buf[:n])

		mutex.Lock()
		clientAddrs[remoteAddr.String()] = remoteAddr
		mutex.Unlock()

		broadcastMessage(message, remoteAddr, conn)
	}
}

func broadcastMessage(message string, sender *net.UDPAddr, conn *net.UDPConn) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, addr := range clientAddrs {
		if addr.String() != sender.String() {
			_, err := conn.WriteToUDP([]byte(message), addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to send message to %s: %s\n", addr, err)
			}
		}
	}
}

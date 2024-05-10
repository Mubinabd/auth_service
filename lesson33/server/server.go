package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

type Client struct {
	conn net.Conn
	name string
}

var (
	clients []*Client
	mutex   sync.Mutex
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listening: %s\n", err)
		os.Exit(1)
	}
	defer ln.Close()
	fmt.Println("TCP server listening on port 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error accepting connection: %s\n", err)
			continue
		}
		client := &Client{conn: conn, name: conn.RemoteAddr().String()}
		mutex.Lock()
		clients = append(clients, client)
		mutex.Unlock()
		fmt.Printf("Client connected: %s\n", client.name)
		go handleConnection(client)
	}
}

func handleConnection(client *Client) {
	defer func() {
		client.conn.Close()
		removeClient(client)
	}()

	scanner := bufio.NewScanner(client.conn)
	for scanner.Scan() {
		text := scanner.Text()
		broadcastMessage(fmt.Sprintf("%s: %s", client.name, text), client)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("%s: %s\n", client.name, err)
	}
}

func broadcastMessage(message string, sender *Client) {
	mutex.Lock()
	defer mutex.Unlock()
	for _, client := range clients {
		if client != sender {
			fmt.Fprintf(client.conn, message+"\n")
		}
	}
}

func removeClient(client *Client) {
	mutex.Lock()
	defer mutex.Unlock()
	for i, c := range clients {
		if c == client {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}

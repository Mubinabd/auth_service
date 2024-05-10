package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting: %s\n", err)
		return
	}
	defer conn.Close()

	// Start a goroutine to read from the server
	go func() {
		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// Send messages to the server
	fmt.Println("START")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Fprintf(conn, "%s\n", scanner.Text())
	}
	fmt.Println("END")
	
}

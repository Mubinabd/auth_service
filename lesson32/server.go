package main

import (
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 1024)
	_, err := conn.Read(b)
	if err != nil {
		fmt.Println("error while reading", err)
		return
	}

	fmt.Println("This message: ", string(b))

	// conn.Write([]byte("We received\n"))
}

func main() {
	
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		fmt.Println("Error resolving address:", err.Error())
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println("Server is listening on 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		handleClient(conn)
	}
}

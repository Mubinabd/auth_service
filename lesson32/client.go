package main

import (
	"fmt"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message := "Hello World, i'm a client"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("error while writing as client", err)
		return
	}

	// serverdan javobni o'qish
	// b := make([]byte, 1024)
	// _, err = conn.Read(b)
	// if err != nil {
	// 	fmt.Println("error while reading as client", err)
	// 	return
	// }

	fmt.Println("message from server")
}

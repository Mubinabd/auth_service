package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func str(wg *sync.WaitGroup, text string) {
	fmt.Println(text)
	time.Sleep(4 * time.Second)
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(4)

	for i := 4; i > 0; i-- {
		go str(&wg, "Hello World "+strconv.Itoa(i))
	}

	wg.Wait()
	fmt.Println("SOMETHING")
}

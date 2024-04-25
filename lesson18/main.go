package main

import "fmt"

//func summer(a int) int {
//	fmt.Println("SUMMER!", a)
//	return a + a
//}

//	func main() {
//		// Go Scheduler -- jadvalchi
//		//for i := 0; i < 10; i++ {
//		//	z := go summer(i)
//		//	fmt.Println(z)
//		//}
//		//time.Sleep(time.Second)
//		//fmt.Println("Goodbye, World!")
//
//		// Channels
//		//var channel = make(chan int)
//		//defer close(channel)
//
//		// arrow operator <-
//
//		// write
//		//channel<- 12432
//
//		// read
//		//a := <- channel
//
//		// unbuffered channel
//		//go square(channel, 43)
//		//time.Sleep(3 * time.Second)
//		//
//		//fmt.Println(<-channel)
//		//fmt.Println(1)
//		//fmt.Println(<-channel)
//		//fmt.Println(2)
//		//fmt.Println(<-channel)
//		//fmt.Println(3)
//		//fmt.Println(<-channel)
//		//fmt.Println(4)
//		//fmt.Println(<-channel)
//
//		// Buffered channel
//		ch := make(chan int, 1)
//		t := 2
//		ch <- t
//		//go cube(ch, t)
//		//close(ch)
//		fmt.Println(<-ch)
//		//fmt.Println(<-ch)
//		//ch <- (t + 1) * (t + 1) * (t + 1)
//		//ch <- (t + 2) * (t + 2) * (t + 2)
//		//ch <- (t + 3) * (t + 3) * (t + 3)
//		////close(ch)
//		//
//		//fmt.Println(len(ch), cap(ch))
//		////time.Sleep(time.Second)
//		//i := 0
//		//for v := range ch {
//		//	fmt.Println(len(ch), cap(ch), i)
//		//
//		//	fmt.Println(v)
//		//	i++
//		//}
//		//fmt.Println(<-ch)
//		//fmt.Println(<-ch)
//	}
//
//	func main() {
//		ch1 := make(chan int)
//		ch2 := make(chan int)
//		go func() {
//			v := 1
//			ch1 <- v
//			v2 := <-ch2
//			fmt.Println(v, v2)
//		}()
//		v := 2
//		ch2 <- v
//		v2 := <-ch1
//		fmt.Println(v, v2)
//	}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2, "f")
	}()
	v := 2
	var v2 int
	select {
	case ch2 <- v:
	case v2 = <-ch1:
	}
	fmt.Println(v, v2)
}

func cube(x chan int, t int) {
	x = make(chan int, 1)
	x <- t * t * t
	//x <- (t + 1) * (t + 1) * (t + 1)
	//x <- (t + 2) * (t + 2) * (t + 2)
	//x <- (t + 3) * (t + 3) * (t + 3)
}

//func square(x chan int, t int) {
//	x <- t * t
//	fmt.Println(1, "s")
//	x <- t * t * t
//	fmt.Println(2, "s")
//	x <- t * t * t
//	fmt.Println(3, "s")
//	x <- t * t * t
//	fmt.Println(4, "s")
//	x <- t * t * t
//	fmt.Println(5, "s")
//}

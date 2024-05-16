package main

import (
	"context"
	"fmt"
	"time"
)

type ourKey string

func main() {

	//t := ourKey("Hello")
	//ctx := context.Background()
	//ctx = context.WithValue(ctx, t, "0.0.1")
	//ctx = context.WithValue(ctx, "key", "val")
	//checker(ctx, t)
	//ctx, cancel := context.WithCancel(context.Background())
	//ctx = context.WithValue(ctx, "version", "1.1.4")
	//go greeting(ctx)
	//time.Sleep(1 * time.Second)
	//cancel()
	//time.Sleep(1 * time.Second)

	//ctx, cancel := context.WithCancel(context.Background())
	//time.AfterFunc(2*time.Second, cancel)
	//
	//sayMyName(ctx)

	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(-2*time.Second))
	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//
	//select {
	//case <-ctx.Done():
	//	defer cancel()
	//	fmt.Println(ctx.Err(), 4324)
	//case <-time.After(3 * time.Second):
	//	fmt.Println(ctx.Err())
	//	fmt.Println("Hello, everything is okay")
	//}
	context.AfterFunc()

}

func sayMyName(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-time.After(3 * time.Second):
		fmt.Println("Hello, everything is okay")
	}
}

func greeting(ctx context.Context) {
	//ctx = context.WithValue(ctx, "key", 4215352345)
	i := 0
	for {
		fmt.Println("Hello World", ctx.Value("version"), ctx.Value("key"), i)
		select {
		case <-ctx.Done():
			i += 1000000000
			goto here
		default:
			break
		}

		fmt.Println(ctx.Err())
		i++

	}
here:
	fmt.Println(i)
	fmt.Println("Hello World", ctx.Value("version"), ctx.Value("key"), i)

}
func checker(ctx context.Context, key ourKey) {
	k, ok := ctx.Value(key).(ourKey)
	if ok {
		fmt.Println("our key is", k)
		return
	}
	fmt.Println("something happened", k, ok)
}

// https://david-yappeter.medium.com/context-in-go-programming-part-1-3a8d470617d0
// https://go.dev/talks/2014/gotham-context.slide#21

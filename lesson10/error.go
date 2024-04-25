package main

import (
	"fmt"
)

func main() {
	// text := "Hello World"
	// fmt.Println(strings.Split(text, " "))

	// strings.

	// ERROR

	var c byte = 'a'

	c, err := lowerToUpper(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(c)

	s := "a"
	if len(s ) > 1 {
		fmt.Println(s[1])
	} else {
		panic("not enough length")
	}
	fmt.Println("fdf")
	
}

func lowerToUpper(a byte) (byte, error) {
	if a >= 'a' && a <= 'z' {
		return a - 32, nil
	}

	return 0, fmt.Errorf("this is not lower letter")
}

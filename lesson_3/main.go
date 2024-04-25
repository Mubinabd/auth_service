package main

import "fmt"

//const ip = "127.1.1.1"
//
//const pi = 3.1415

func main() {
	//const b = 54
	// %
	//fmt.Println(fmt.Sprintf("bu yerda text bor %d", b))

	//fmt.Println(5345 > 654)
	//fmt.Println(5345 > 654 && 1 == 1)
	//fmt.Println(!true)
	// ! =   ==  >= <= > <
	// && ||   !
	//a := 54
	//if a > 50 {
	//	fmt.Println("a katta 50")
	//} else if a == 50 {
	//	fmt.Println("a kichkina 50dan")
	//}

	a := 54

	switch {
	case a == 54:
		fmt.Println("a 54")
		fallthrough
	case a > 53:
		fmt.Println("a 53")
		fallthrough
	case a < 52:
		fmt.Println("a 53")
		fallthrough
	default:
		fmt.Println("default")
	}

}

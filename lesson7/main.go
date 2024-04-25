package main

import "fmt"

//func main() {
//
//	q := max
//
//	e := func(a, b int) int {
//		if a > b {
//			return a
//		}
//		return b
//	}
//
//	fmt.Println(max(43, 543), q(4, 35))
//	fmt.Println(e(1, 2))
//}

//func max(a int, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//func main() {
//	//var a = 34567
//	////
//	//fmt.Println(a, &a, *&a, &a)
//
//	//var n = 76
//	//var p *int
//	//p = &n
//
//	//fmt.Println(*&p)
//
//	// new
//	//p := new([]int)
//	//fmt.Println(p)
//	var a = 23
//	sqr(&a)
//	fmt.Println(a)
//}
//
//func sqr(b *int) {
//	*b = *b * *b
//}

//func failedUpdate(g *int) {
//	x := 10
//	g = &x
//}
//
//func main() {
//	var f *int // f is nil
//	failedUpdate(f)
//	fmt.Println(f) // prints nil
//}

//	func failedUpdate(px *int) {
//		x2 := 20
//		px = &x2
//	}
//
//	func update(px *int) {
//		*px = 20
//	}
//
//	func main() {
//		x := 10
//		failedUpdate(&x)
//		fmt.Println(x) // prints 10
//		update(&x)
//		fmt.Println(x) // prints 20
//	}
func main() {
	//var a = [4]int{1, 2, 3, 4}
	// array with pointer
	//fmt.Printf("%p %p %p %p", a, &a, &a[0], &a[1])

	//var a = []int{1, 2}
	////// array with slice
	//fmt.Printf("%p %p %p %p\n", &a, a, &a[0], &a[1])
	//a = append(a, 3)
	//a = append(a, 4)
	//fmt.Printf("%p %p %p %p", &a, a, &a[0], &a[1])

	//var a [3]*int
	//var b = [3]int{3, 4, 5}
	//a[0] = &b[0]
	//fmt.Println(*a[0])
	//var a = new([]int)
	//*a = append(*a, 23)
	//fmt.Println(a)

	//func sqr(a *[5]int) {
	//	for i, v := range a {
	//		a[i] = v * v
	//	}
	var arr = make([]int, 3, 5)
	//arr := []int{0, 1, 2, 3, 4}

	doubleInts(&arr)

	fmt.Println(arr)
}

func doubleInts(array *[]int) {
	for i := 0; i < len(*array); i++ {
		(*array)[i] = (*array)[i] + 2
	}
	*array = append(*array, 3)
}

package main

import "fmt"

func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	n -= 1

	for i, j := 0, n; n >= 0; n-- {
		if nums[i] < 0 && -nums[i] >= nums[j] {
			result[n] = nums[i] * nums[i]
			i++
		} else {
			result[n] = nums[j] * nums[j]
			j--
		}
	}

	return result
}

//func main() {
//
//	//text := "Салом Дунё 你好呀"
//	//
//	//for _, v := range text {
//	//
//	//}
//	//
//	//next:
//	//	fmt.Println("Shettaman")
//	t, _ := addAndMul(54, 5)
//	fmt.Println(t)
//}
//

//func main() {
//
//	a := calc(10, "-", 5)
//	fmt.Print(a)
//}

//func calc(a int, b string, c int) int {
//
//	if string(b) == "-" {
//		return a - c
//	} else if string(b) == "+" {
//		return a + c
//	} else if string(b) == "/" {
//		return a / c
//	} else if string(b) == "*" {
//		return a * c
//	}
//
//	//z := string([]byte{'a', 'b', 'b'})
//	return 0
//}

//func main() {
//	result := removeVowels("Hello World")
//	println(result)
//}
//
//func removeVowels(text string) string {
//	txt := ""
//	for _, v := range text {
//		if !isVowel(byte(v)) {
//			txt += string(v)
//		}
//	}
//	return txt
//}
//
//func isVowel(char byte) bool {
//	switch char {
//	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
//		return true
//	}
//	return false
//}

//func add(a, b int) int { return a + b }
//func mul(a, b int) int { return a * b }
//func sub(a, b int) int { return a - b }
//func div(a, b int) int { return a / b }
//
//func main() {
//	var calc []func(int, int) int
//	calc = append(calc, add, mul, sub, div)
//
//	a, b, c := 123, 43, '/'
//
//	switch c {
//	case '+':
//		a = calc[0](a, b)
//	case '*':
//		a = calc[1](a, b)
//	case '-':
//		a = calc[2](a, b)
//	case '/':
//		a = calc[3](a, b)
//	}
//	fmt.Println(a)
//
//	fmt.Println(sum(324, 324, 423, 4, 234, 234))
//}
//
//// variadic parameter/ variable
//func sum(a ...int) (s int) {
//	for _, v := range a {
//		s += v
//	}
//
//	return 65
//}

func main() {
	// anonymous function
	//byteToString := func(text []byte) string {
	//	return string(text)
	//}([]byte{'H', 'e', 'l', 'l', 'o'})
	//
	//fmt.Println(byteToString)
	result := removeVowels("Hello World")

	fmt.Println(result())
}

func removeVowels(text string) func() string {
	txt := []byte{}
	for _, v := range text {
		if !isVowel(byte(v)) {
			txt = append(txt, byte(v))
		}
	}
	return func() string {
		return string(txt)
	}
}

func isVowel(char byte) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

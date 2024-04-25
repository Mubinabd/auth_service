package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Working with Files

	// f, err := os.Open("file.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// // _, err = f.Write([]byte("Hello World"))
	// // if err != nil {
	// // 	panic(err)
	// // }

	// t := make([]byte, 20)
	// _, err = f.Read(t)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(t))

	// s, err := f.Seek(5, 0)
	// if err != nil {
	// 	panic(err)
	// }

	// b2 := make([]byte, 12)
	// _, err = f.Read(b2)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(s, string(b2))

	// _, err = f.Seek(6, 0)
	// if err != nil {
	// 	panic(err)
	// }
	// b3 := make([]byte, 10)
	// n3, err := io.ReadAtLeast(f, b3, 5)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%d bytes @ %s\n", n3, string(b3))

	// st, err := os.Stat("main.go")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%+v", st.ModTime())

	// f, err := os.Create("data.txt")

	// if err != nil {

	// 	panic(err)
	// }

	// defer f.Close()

	// words := []string{"sky", "falcon", "rock", "hawk"}

	// for _, word := range words {

	// 	_, err := f.WriteString(word + "\n")

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// fmt.Println("done")

	// f, err := os.OpenFile("words.txt",os.O_RDWR, 0644)

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// if _, err := f.WriteString("cloud\n"); err != nil {
	// 	panic(err)
	// }
	// _, err = f.Seek(10,0)
	// if err != nil {
	// 	panic(err)
	// }

	// b := make([]byte, 30)
	// _, err = f.Read(b)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(b))

	var files []string

	root := "/Users/husanmusa/go/src/github.com/husanmusa/NT_Golang_10/lesson_3"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return nil
		}

		if info.IsDir() && filepath.Ext(path) == ".go" {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

}

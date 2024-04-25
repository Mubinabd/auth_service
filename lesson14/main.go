package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format("02/01 15:04:05 2006"))
	fmt.Println(t.Format(time.DateOnly))
	fmt.Println(t.Format("02-01-2006 15:04:05"))
	fmt.Println(time.Local)

	// t1 := t.Format("02-01-2006 15:04:05")
	
	// t, err := time.Parse("02-01-2006 15:04:05", t1)
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Printf("%v %T\n", t, t)
	fmt.Println(t.Date())
	fmt.Println(t.Year(), t.Month())
	e, _ := time.LoadLocation("Europe/London")
	fmt.Println(e)
	fmt.Println(t.In(e))
	// fmt.Println(t.Unix())
	fmt.Println(t.Add(100*time.Minute))
	time.Sleep(3*time.Second)
	fmt.Println(t.AddDate(3, 6, 35))
}

package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	t1 := t.AddDate(0, 0, -1)

	fmt.Println(t1.After(t))

	fmt.Println(t.Sub(t1))

	wDay := make(map[int]time.Weekday)
	mDay := make(map[int]time.Month)

	for i := 1; i <= 7; i++ {
		wDay[i] = time.Weekday(i)
	}
	for i := 1; i <= 12; i++ {
		mDay[i] = time.Month(i)
	}
	fmt.Println(time.Weekday(0))
	fmt.Println(mDay[8])
	fmt.Println(wDay[2])

}

// map
// m["Seshanba"] -> 16.04.2024
// m["Yanvar"] -> 01.01.2024

// -------------------------- 2
// map
// wDay[1] -> Tuesday map[int]time.WeekDay
// mDay[3] -> March map[int]time.Month

package main

import (
	"fmt"
	"sort"
)

type intSlice []int

func (i intSlice) Len() int { return len(i) }

func (i intSlice) Less(l, k int) bool {
	return i[k] < i[l]
}

func (i intSlice) Swap(q, e int) {
	i[q], i[e] = i[e], i[q]
}

func main() {
	slc := intSlice{43, 25, 4356, 46, 54, 645, 654, 65, 4}
	sort.Sort(slc)

	fmt.Println(slc)
}

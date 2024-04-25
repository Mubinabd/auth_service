package main

import (
	"fmt"
	lm "lesson11/math"
	"lesson11/math/mul"
	drrr "lesson11/strconv/doubletostring"
	"golang.org/x/exp/slog"
)

func main() {
	fmt.Println(lm.Sqr(54))

	fmt.Println(lm.Cube(4))

	fmt.Println(mul.Double(12))

	fmt.Println(drrr.DtoS(3))

	slog.Info("Hello world")

}

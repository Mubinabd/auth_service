package main

import "fmt"

type CarModel struct {
	brand   string
	year    int
	price   int
	isRun   bool
	battery int
}

type Tesla struct {
	CarModel
	isAutoPilot bool
}

func (c *CarModel) start() {
	if c.battery == 0 {
		fmt.Println("battery is empty")
		return
	}
	c.isRun = true
}

func (c *CarModel) stop() {
	c.isRun = false
}

func (c *CarModel) charging() {
	c.battery = 100
}

func (c *CarModel) run() {
	if !c.isRun {
		fmt.Println("car didn't start")
		return
	}

	fmt.Println("car is running")
	c.battery -= 10
}

func (t *Tesla) start() {
	t.isAutoPilot = true
}

func main() {
	byd := CarModel{
		"BYD",
		2022,
		20000,
		false,
		45,
	}

	t := Tesla{
		byd,
		true,
	}

	t.start()
	fmt.Println(t)

	byd.start()
	fmt.Println(byd)

	byd.stop()

	fmt.Println(byd)

	byd.run()
}

// airplane - methods - fly, arrive
// airport - getReport
// airplane biror airoportdan uchib ketganida airportdagi samolyot kamayadi. Qo'nsa ko'payadi.

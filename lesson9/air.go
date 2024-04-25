package main

import (
	"fmt"
)

type Airplane struct {
	Name   string
	moduli string
	countr string
	isFly  bool
}

type Airport struct {
	name   string
	planes []*Airplane
}

func (plane *Airplane) fly() {
	plane.isFly = true
}

func (plane *Airplane) arrive(port *Airport) {
	plane.isFly = false
	port.takePlane(plane)
}

func (air *Airport) takePlane(plane *Airplane) {
	air.planes = append(air.planes, plane)
}

func (air *Airport) getReport() {
	fmt.Println(len(air.planes))
	fmt.Println(air.name, air.planes)
}

func main() {

	var airplane = &Airplane{
		"Jet",
		"Boeing 747",
		"USA",
		false,
	}
	airplane.fly()
	fmt.Println(airplane)

	var drom = &Airport{
		name: "TOSHKENT",
	}
	airplane.arrive(drom)
	fmt.Println(airplane)
	drom.getReport()
}

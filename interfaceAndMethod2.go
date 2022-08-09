package main

import (
	"fmt"
)

type rectangle interface {
	perimeterOfR()
	areaOfR()
}

type sides struct {
	length  int
	breadth int
}

func (v sides) perimeterOfR() int {
	return 2*v.length + 2*v.breadth
}

func (v sides) areaOfR() int {
	return v.length * v.breadth
}

func main() {
	v := sides{3, 4}
	fmt.Println("Interface :")
	fmt.Println("Perimeter of rectangle :", v.perimeterOfR())
	fmt.Println("Area of rectangle :", v.areaOfR())
}

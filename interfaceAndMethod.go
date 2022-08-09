package main

import (
	"fmt"
	"math"
)

type vertex struct {
	x, y float64
}

func (v vertex) sqrtTest() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v *vertex) Scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func Scale(v vertex, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := vertex{3, 4}
	fmt.Print("Method :")
	fmt.Println(v.sqrtTest())
	fmt.Print("Pointer Receivers :")
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.sqrtTest())
	fmt.Print("Pointer And Functions :")
	Scale(v, 10)
	fmt.Println(v.sqrtTest())
}

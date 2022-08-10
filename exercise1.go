package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	fmt.Println("X value is :", x)
	z := math.Sqrt(x)
	fmt.Println("Z value before return :", z)
	z -= (z*z - x) / (2 * z)
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

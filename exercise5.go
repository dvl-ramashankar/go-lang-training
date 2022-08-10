package main

import (
	"fmt"
	"math"
)

type MyError struct {
	err string
}

func (e *MyError) Error() string {
	return fmt.Sprintf(e.err)
}

func run() error {
	return &MyError{"The given number is negative"}
}

func Sqrt(x float64) (float64, error) {
	if x > 0 {
		return math.Sqrt(x), nil
	}
	return x, run()
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(25))
	fmt.Println(Sqrt(-25))
}

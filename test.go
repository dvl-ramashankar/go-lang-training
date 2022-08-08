package main

import (
	"fmt"
	"log"
)

var c, python, java bool
var i int
var a int
var f float64
var b bool
var s string

func constant() int {
	const num3 = 10
	return num3
}

func sum() int {
	var num1, num2 = 5, 5
	num4 := 10
	var result int
	result = num1 + num2 + constant() + num4
	return result
}
func stmt() string {
	return "Hello, RK"
}

func loop(x int, y int) {
	log.Println("Executing Loop")
	for i := x; i <= y; i++ {
		log.Println("Hello")
	}
	log.Println("Loop Completed")
}

func calci(x int, y int) (int, int, int, int, int) {
	sub := x - y
	add := x + y
	div := x / y
	mult := x * y
	mod := x % y
	return sub, add, div, mult, mod
}

func calciResult() {
	sub, add, div, mult, mod := calci(15, 2)
	log.Println("Subtraction : ", sub)
	log.Println("Addition : ", add)
	log.Println("Division : ", div)
	log.Println("Multiplication : ", mult)
	log.Println("Modulus : ", mod)
}
func array() {
	arr := [4]string{"rk", "rk1"}
	fmt.Println("Elements of the array:")
	for i := 0; i < 2; i++ {
		fmt.Println(arr[i])
	}
	log.Println("Completed Array Method")
}

func test(x, y int) (a, b int) {
	a = x * 2
	b = y * 2
	return
}

func main() {
	fmt.Println(stmt())
	log.Print("Addition :")
	fmt.Println(sum())
	log.Print("Constant :", constant())
	loop(1, 4)
	log.Print("Subtraction, Addition, Division, Multiplication, Modulus :")
	fmt.Println(calci(15, 2))
	calciResult()
	array()
	log.Print("Test return of muliplte value :")
	fmt.Println(test(2, 4))
	fmt.Println("The package Variable : ", i, c, python, java)
	fmt.Println("The dafault values are : ", i, f, b, s)
}

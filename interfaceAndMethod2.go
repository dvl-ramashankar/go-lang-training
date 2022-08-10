package main

import (
	"fmt"

	l "demo/exportPackage"
	d "demo/packageTest"
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

type point struct {
	l int
	b int
}

func (k *point) areaOfR() int {
	return k.l * k.b
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	v := sides{3, 4}
	fmt.Println("Interface :")
	fmt.Println("Perimeter of rectangle :", v.perimeterOfR())
	fmt.Println("Area of rectangle :", v.areaOfR())
	k := &sides{7, 5}
	fmt.Println("Area of rectangle Pointer :", k.areaOfR())
	fmt.Print("Empty Interface :")
	var i interface{}
	describe(i)
	i = "hello"
	describe(i)
	fmt.Print("Type Switch :")
	do(21)
	do("hello")
	do(true)
	l.M1()
	l.M2()
	fmt.Println("The addition in different package : ", d.M3(5, 5))
}

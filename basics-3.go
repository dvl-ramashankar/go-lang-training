package main

import (
	"fmt"
	"math"
	"strings"
)

func pointerTest() {
	fmt.Println("Pointer :")
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i
	p = &j          // point to j
	*p = *p / 37    // divide j through the pointer
	fmt.Println(j)  // see the new value of j
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}       // has type Vertex
	v2 = Vertex{X: 6, Y: 5} // Y:0 is implicit
	v3 = Vertex{Y: 8}       // X:0
	v4 = Vertex{}           // X:0 and Y:0
	p  = &Vertex{3, 4}      // has type *Vertex
)

func arrayTest() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("Array : ", a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Prime No :", primes)
}

func sliceTest() {
	no := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
	var s []int = no[1:8]
	fmt.Println("Slice : ", s)
}

func sliceTest2() {
	fmt.Print("Slice :")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceDefaultAndCapacity() {
	fmt.Print("Slice Default :")
	s := []int{2, 3, 5, 7, 11, 13}
	var x []int = s[1:4]
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
	var y []int = s[:2] // zero for the low bound
	fmt.Printf("len=%d cap=%d %v\n", len(y), cap(y), y)
	var z []int = s[1:] //length of the array for the high bound
	fmt.Printf("len=%d cap=%d %v\n", len(z), cap(z), z)
	y = y[:0]
	fmt.Printf("len=%d cap=%d %v\n", len(y), cap(y), y)
}

func sliceWithMake() {
	fmt.Print("Slice with Make :")
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func sliceOfSlice() {
	fmt.Println("Slice of Slice :")
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendSlice() {
	fmt.Println("Slice Append :")
	var s []int
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// append works on nil slices.
	s = append(s, 0)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// The slice grows as needed.
	s = append(s, 1)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	// We can add more than one element at a time.
	s = append(s, 1, 3, 4)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func rangeTest() {
	fmt.Println("Range :")
	r := []int{1, 2, 4, 8, 16}
	for i, v := range r {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func mapTest() {
	fmt.Println("Map :")
	type Vertex struct {
		Lat, Long float64
	}
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"":          {0.23, 0.56},
		"Google":    {23.34, 56.54},
	}
	fmt.Println(m)
}

func mutatingMap() {
	fmt.Println("Mutating Map :")
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	m["Answer"] = 40
	s, ok := m["Answer"]
	fmt.Println("The value:", s, "Present?", ok)
}

func functionValues() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(10, 0))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	pointerTest()
	fmt.Println("Structure :", v1, v2, v3, v4, p)
	arrayTest()
	sliceTest()
	sliceTest2()
	sliceDefaultAndCapacity()
	sliceWithMake()
	sliceOfSlice()
	appendSlice()
	rangeTest()
	mapTest()
	mutatingMap()
	functionValues()
}

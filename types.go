package main

import "fmt"

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
func main() {
	pointerTest()
	fmt.Println("Structure :", v1, v2, v3, v4, p)
	arrayTest()
	sliceTest()
	sliceTest2()
	sliceDefaultAndCapacity()
}

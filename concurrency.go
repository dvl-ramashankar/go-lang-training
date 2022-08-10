package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}

func testThread(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(s)
	}
}
func sumTest(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func channelTest() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sumTest(s[:len(s)/2], c)
	go sumTest(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, " + ", y, " = ", x+y)
}

func bufferedChannelTest() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Print(" ", <-ch)
	fmt.Println("	", <-ch)
}

func fibonacciUsingChannels(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func fibonacciUsingChannelsCall() {
	c := make(chan int, 10)
	go fibonacciUsingChannels(cap(c), c)
	for i := range c {
		fmt.Print(" ", i)
	}
}

func main() {
	go hello()
	fmt.Println("main function")
	go testThread("Hi")
	testThread("All!!")
	time.Sleep(time.Second)
	fmt.Print("Channels =  ")
	channelTest()
	fmt.Print("Buffered Channels =  ")
	bufferedChannelTest()
	fmt.Print("Range And Close =  ")
	fibonacciUsingChannelsCall()
}

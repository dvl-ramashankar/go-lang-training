package main

import (
	"fmt"
	"math"
	"time"
)

func loopTest(x, z int) (y int) {
	for ; x <= z; x++ {
		y += x
	}
	return
}

func while(x int) int {
	for x > 10 {
		x += x
	}
	return x
}

func iftest(x, y int) string {
	if x < y {
		return "The x is smaller than y "
	} else if x == y {
		return "The x is equal to y"
	} else {
		return "The x is greater than y"
	}
}

func pow(x, y, z float64) float64 {
	if v := math.Pow(x, y); v < z {
		return v
	}
	return z
}

func switchTest(str string) {
	fmt.Println("Switch-case :")
	switch str {
	case "Monday":
		fmt.Println("Monday!")
	case "Tuesday":
		fmt.Println("Tuesday!")
	case "Wednesday":
		fmt.Println("Wednesday!")
	case "Thursday":
		fmt.Println("Thursday!")
	case "Friday":
		fmt.Println("Friday!")
	case "Saturday":
		fmt.Println("Saturday!")
	case "Sunday":
		fmt.Println("Sunday!")
	default:
		fmt.Println("Invalid weekday!")
	}
}

func switchTest2() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferTest(x, y int) {
	fmt.Println("Defer :")
	for ; x <= y; x++ {
		defer fmt.Println(x)
	}
	fmt.Println("Executed defer statement.")
}

func main() {
	fmt.Println("The For Loop :", loopTest(5, 10))
	fmt.Println("The While Loop :", while(5))
	fmt.Println("If-else :", iftest(15, 11))
	fmt.Println("Math-Pow : ", pow(4, 2, 20), pow(4, 3, 20))
	switchTest(time.Now().Weekday().String())
	switchTest2()
	deferTest(1, 2)
}

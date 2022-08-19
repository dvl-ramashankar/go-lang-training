package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	arr := strings.Fields(s)
	fmt.Println(len(arr))
	fmt.Println(arr)
	x := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				x++
			}
		}
		m[arr[i]] = x
		x = 0
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

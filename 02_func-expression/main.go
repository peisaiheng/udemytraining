package main

import "fmt"

func wrapper(first int) func(second int) int {
	x := first
	return func(second int) int {
		return x * second
	}
}

func main() {
	result := wrapper(99)(2)
	fmt.Println(result)
}

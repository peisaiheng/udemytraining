package main

import "fmt"

func varia(x ...int32) int32 {
	var highest int32
	for _, value := range x {
		if value > highest {
			highest = value
		}
	}
	return highest
}

func main() {
	list := []int32{32, 5, 10, 121, 58}
	fmt.Println(varia(list...))
}

package main

import "fmt"

func half(x int) (float64, bool) {
	return float64(x) / 2, x%2 == 0
}

func main() {
	h, even := half(5)
	fmt.Println(h, even)
}

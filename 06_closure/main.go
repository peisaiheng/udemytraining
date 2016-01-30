package main

import "fmt"

func main() {
	makeEvenGenerator()()
	makeEvenGenerator()()
	makeEvenGenerator()()

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}

func makeEvenGenerator() func() int {
	i := 0
	return func() int {
		i += 2
		return i
	}
}

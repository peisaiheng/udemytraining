package main

import "fmt"

func main() {
	var something *[]string = new([]string)
	func(s *[]string) {
		*s = append(*s, "golang")
	}(something)
	fmt.Println(something)
	fmt.Printf("%T\n", something)

}

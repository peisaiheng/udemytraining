package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	hello := "Hello"
	//var userInput string

	fmt.Print("Please enter your name:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Printf("%v %v", hello, scanner.Text())
	//ExampleScanner_lines()

}

// The simplest use of a Scanner, to read standard input as a set of lines.
func ExampleScanner_lines() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

package main

import "fmt"

type ToDo struct {
	ID    float32
	Name  string
	Email string
}

func main() {
	v := make([]ToDo, 0)
	//	for i := 0; i < 25; i++ {
	//		if i >= len(v) {
	//			v = append(v, i)
	//		}
	//		v[i] = i+1
	//		fmt.Println(v)
	//	}
	fmt.Println(v)

}

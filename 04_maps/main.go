package main

import "fmt"

func main() {

	var myGreeting = make(map[string]map[string]string)
	myGreeting["en"] = map[string]string{
		"formal":   "Greetings",
		"informal": "Hey",
	}
	myGreeting["es"] = map[string]string{
		"formal":   "Saludos",
		"informal": "Hola",
	}

	myGreeting2 := map[string]map[string]string{
		"en": map[string]string{
			"formal":   "Greetings",
			"informal": "Hey",
		},
		"es": map[string]string{
			"formal":   "Saludos",
			"informal": "Hola",
		},
	}

	fmt.Println(myGreeting["es"]["informal"])
	fmt.Println(myGreeting)
	fmt.Println(myGreeting2["en"]["formal"])
	fmt.Println(len(myGreeting))
}

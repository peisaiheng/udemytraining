package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	name := "stiff"

	//	parse template
	tpl, err := template.ParseFiles("name.gohtml")

	// execute template
	err = tpl.Execute(os.Stdout, name)
	if err != nil {
		log.Fatalln(err)
	}
}

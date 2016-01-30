package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Page struct {
	Title string
	Body  string
}

func main() {
	var tpl *template.Template
	tpl, err := tpl.ParseGlob("tmp/*.go.htm")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, Page{
		Title: "Execute",
		Body:  "Template was not stated, default to first template",
	})
	fmt.Printf("\n******************************\n")

	err = tpl.ExecuteTemplate(os.Stdout, "index.go.htm", Page{
		Title: "ExecuteTemplate",
		Body:  "Template is: index.go.htm",
	})
	fmt.Printf("\n******************************\n")

	err = tpl.ExecuteTemplate(os.Stdout, "home.go.htm", Page{
		Title: "ExecuteTemplate",
		Body:  "Template is: home.go.htm",
	})
	fmt.Printf("\n******************************\n")
}

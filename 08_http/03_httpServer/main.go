package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func upTown(res http.ResponseWriter, req *http.Request) {
	fmt.Println(res, req)
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	var dogName string
	fs := strings.Split(req.URL.Path, "/")
	if len(fs) >= 3 {
		dogName = fs[2]
	}

	io.WriteString(res, `
	Dog name is: <strong>`+dogName+`</strong><br>
	<img src="/assets/stiff.jpg" width="900px" width="auto">`)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog/", upTown)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

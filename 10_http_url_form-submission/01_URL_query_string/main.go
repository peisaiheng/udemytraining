package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "q"
		value := req.URL.Query().Get(key)
		io.WriteString(res, value)
	})
	http.ListenAndServe(":8080", nil)
}

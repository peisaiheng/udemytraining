package main

import "net/http"

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "q"
		value := req.FormValue(key)

	})
	http.ListenAndServe(":8080", nil)
}

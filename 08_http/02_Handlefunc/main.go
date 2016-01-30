package main

import (
	"io"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/dog/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, `<img src="http://bebusinessed.com/wp-content/uploads/2014/03/734899052_13956580111.jpg">`)
	})
	mux.HandleFunc("/cat/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		io.WriteString(res, `<img src="http://scienceblogs.com/gregladen/files/2012/12/Beautifull-cat-cats-14749885-1600-1200.jpg">`)
	})

	http.ListenAndServe(":8080", mux)

}

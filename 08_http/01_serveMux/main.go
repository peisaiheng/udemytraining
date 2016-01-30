package main

import (
	"io"
	"net/http"
)

type Hdog int
type Hcat int

func (Hdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="http://bebusinessed.com/wp-content/uploads/2014/03/734899052_13956580111.jpg">`)
}

func (Hcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="http://scienceblogs.com/gregladen/files/2012/12/Beautifull-cat-cats-14749885-1600-1200.jpg">`)
}

func main() {
	var dog Hdog
	var cat Hcat

	mux := http.NewServeMux()
	mux.Handle("/dog/", dog)
	mux.Handle("/cat/", cat)

	http.ListenAndServe(":8080", mux)

}

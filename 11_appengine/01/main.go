package main

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
)

func init() {
	http.HandleFunc("/", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	// New context (only used if serving from app engine)
	ctx := appengine.NewContext(req)

	item1 := memcache.Item{
		Key:   "foo",
		Value: []byte("bar"),
	}

	memcache.Set(ctx, &item1)

	item, err := memcache.Get(ctx, "foo")
	if err != nil && err != memcache.ErrCacheMiss {
		fmt.Fprintln(res, err)
	}
	if err == nil {
		fmt.Fprintln(res, item)
	} else {
		fmt.Fprintln(res, "memcache miss")
	}

}

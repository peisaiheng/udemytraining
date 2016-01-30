package main

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
)

func getSession(req *http.Request) (*memcache.Item, error) {
	ctx := appengine.NewContext(req)

	// Find if user already has session from cookies
	cookie, err := req.Cookie("session")
	if err != nil {
		// No cookie found return empty memcache item
		return &memcache.Item{}, err
	}
	// Cookie found, get session using cookie value
	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil {
		// No  session found,
		// return empty memcache item
		return &memcache.Item{}, err
	}
	// Return session in memcache
	return item, nil
}

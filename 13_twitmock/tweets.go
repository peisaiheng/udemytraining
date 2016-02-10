package main

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

func getTweets(req *http.Request, user *User) ([]Tweet, error) {
	ctx := appengine.NewContext(req)

	var tweets []Tweet
	q := datastore.NewQuery("Tweets")

	// If user is not NIL
	if user != nil {
		// show tweets of a specific user
		userKey := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
		q = q.Ancestor(userKey)
	}

	// Get all tweets
	q = q.Order("-Time").Limit(20)
	_, err := q.GetAll(ctx, &tweets)
	return tweets, err
}

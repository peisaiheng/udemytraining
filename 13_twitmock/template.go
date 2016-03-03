package main

import (
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
)

func serveTemplate(res http.ResponseWriter, req *http.Request, templateName string) {
	ctx := appengine.NewContext(req)
	memItem, err := getSession(req)

	// USER IS NOT LOGGED IN

	if err != nil {
		// Redirect user if they access tweet page
		if templateName == "tweets.html" {
			http.Redirect(res, req, "/", 302)
			return
		}
		// Execute template with LoginFail is true
		if templateName ==  "signin.html"{
			tpl.ExecuteTemplate(res, templateName, &SessionData{LoginFail:true})
			return
		}
		// execute template with empty session data
		tpl.ExecuteTemplate(res, templateName, &SessionData{})
		return
	}

	// USER IS LOGGED IN

	// Redirect user to home page
	// if user goes to signup page
	if templateName == "signup.html" || templateName == "signin.html" {
		http.Redirect(res, req, "/tweets", 302)
	}

	// Initialize session data with user info
	var sd SessionData
	json.Unmarshal(memItem.Value, &sd)
	sd.LoggedIn = true

	// Initialize session data with tweets info
	tweets, err := getTweets(req, nil)
	if err != nil {
		log.Errorf(ctx, "error getting tweets: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	sd.Tweets = tweets

	// Execute template with data
	tpl.ExecuteTemplate(res, templateName, &sd)
}

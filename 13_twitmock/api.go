package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"io/ioutil"
	"net/http"
	"time"
)

func checkUsername(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	bs, err := ioutil.ReadAll(req.Body)
	sbs := string(bs)
	log.Infof(ctx, "REQUEST BODY: %v", sbs)
	var user User
	key := datastore.NewKey(ctx, "Users", sbs, 0, nil)
	err = datastore.Get(ctx, key, &user)
	// if there is an err, there is NO user
	log.Infof(ctx, "ERR: %v", err)
	if err != nil {
		// there is an err, there is a NO user
		fmt.Fprint(res, "false")
		return
	}
	fmt.Fprint(res, "true")
}

func createUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	// Creating a hashed pw to be stored in db
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		log.Errorf(ctx, "Error creating hpassword: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	user := User{
		Email:    req.FormValue("email"),
		UserName: req.FormValue("username"),
		Password: string(hashedPass),
	}
	key := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
	key, err = datastore.Put(ctx, key, &user)
	if err != nil {
		log.Errorf(ctx, "error adding todo: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}

	createSession(res, req, user)
	http.Redirect(res, req, "/tweets", 302)
}

func createSession(res http.ResponseWriter, req *http.Request, user User) {
	ctx := appengine.NewContext(req)
	// Creating a new UUID for new session
	id, err := uuid.NewV4()
	// Setting new cookie with UUID as value
	cookie := &http.Cookie{
		Name:  "session",
		Value: id.String(),
		Path:  "/",
		//		UNCOMMENT WHEN DEPLOYED:
		//		Secure: true,
		//		HttpOnly: true,
	}
	http.SetCookie(res, cookie)

	// Set MEMCAHE session data (sd)
	json, err := json.Marshal(user)
	if err != nil {
		log.Errorf(ctx, "Error marshalling during user creation: %V", err)
		http.Error(res, err.Error(), 500)
		return
	}

	// Set memcache
	sd := memcache.Item{
		Key:   id.String(),
		Value: json,
	}
	memcache.Set(ctx, &sd)
}

func loginProcess(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	username := req.FormValue("username")

	// Set up user dst
	var user User

	key := datastore.NewKey(ctx, "Users", username, 0, nil)
	err := datastore.Get(ctx, key, &user)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.FormValue("password"))) != nil {
		// Fail to log in
		var sd SessionData
		sd.LoginFail = true
		http.Redirect(res, req, "/", 302)
	}
	user.UserName = username
	// success at logging in
	createSession(res, req, user)
	// Redirect to homepage
	http.Redirect(res, req, "/tweets", 302)
}

func logout(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	// Check for session cookie
	cookie, err := req.Cookie("session")
	// cookie is not set
	if err != nil {
		http.Redirect(res, req, "/", 302)
		return
	}

	//clear MEMCACHE and COOKIE if set
	sd := memcache.Item{
		Key:        cookie.Value,
		Value:      []byte(""),
		Expiration: time.Duration(1 * time.Microsecond),
	}
	memcache.Set(ctx, &sd)

	cookie.MaxAge = -1
	http.SetCookie(res, cookie)

	// redirect
	http.Redirect(res, req, "/", 302)
}

func postTweet(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)

	// Check to see if user is logged in
	// Reject post if user is not logged in
	memItem, err := getSession(req)
	if err != nil {
		log.Infof(ctx, "Attempt to tweet from logged out user")
		http.Error(res, "You must be logged in.", http.StatusForbidden)
		return
	}

	// Declare variable of type user
	// Initialize with values from MEMCACHE item
	var user User
	json.Unmarshal(memItem.Value, &user)

	// Declare variable of type tweet
	// Initialize with form values
	tweet := Tweet{
		Msg:      req.FormValue("tweet"),
		Time:     time.Now(),
		UserName: user.UserName,
	}

	// Store tweet in datastore
	// using incomplete key with user as parent
	userKey := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
	tweetKey := datastore.NewIncompleteKey(ctx, "Tweets", userKey)
	ukey, err := datastore.Put(ctx, tweetKey, &tweet)
	if err != nil {
		log.Errorf(ctx, "error adding todo: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}
	var newTweet Tweet
	err = datastore.Get(ctx, ukey, newTweet)

	// Redirect user back to home page
	http.Redirect(res, req, "/tweets", 302)
}

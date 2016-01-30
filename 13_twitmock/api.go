package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
	"io/ioutil"
	"net/http"
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
	} else {
		fmt.Fprint(res, "true")
	}
}

func createUser(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	user := User{
		Email:    req.FormValue("email"),
		UserName: req.FormValue("username"),
		Password: req.FormValue("password"),
	}
	key := datastore.NewKey(ctx, "Users", user.UserName, 0, nil)
	key, err := datastore.Put(ctx, key, &user)
	if err != nil {
		log.Errorf(ctx, "error adding todo: %v", err)
		http.Error(res, err.Error(), 500)
		return
	}

	createSession(res, req, user)
	http.Redirect(res, req, "/", 302)
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
		Key:        id.String(),
		Value:      json,
	}
	memcache.Set(ctx, &sd)
}

func loginProcess(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	ctx := appengine.NewContext(req)
	username := req.FormValue("username")
	pw := req.FormValue("password")

	// Set up user dst
	var user User

	key := datastore.NewKey(ctx, "Users", username, 0, nil)
	err := datastore.Get(ctx, key, &user)
	if err != nil || pw != user.Password {
		// Fail to log in
		var sd SessionData
		sd.LoginFail = true
		tpl.ExecuteTemplate(res, "index.html", sd)
		return
	} else {
		user.UserName = username
		// success at loggin in
		createSession(res, req, user)
		// Redirect to homepage
		http.Redirect(res, req, "/", 302)
	}
}
//func logout(res http.Response, req *http.Request, _ httprouter.Params) {
//
//}

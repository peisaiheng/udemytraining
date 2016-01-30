package main

import (
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/memcache"
	"html/template"
	"net/http"
	"strings"
)

var tpl *template.Template

type Animal struct {
	Term       string
	Definition string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("/", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	// Handle user session
	user := handleUserSession(res, req) // returns map of memcache values
	username := user["name"]

	//Handle POST
	if req.Method == "POST" {
		handlePost(res, req)
	}

	var animalList []Animal

	// Handle GET request
	if req.Method == "GET" {

		// Checking URL path for term
		if req.URL.Path != "/" {
			term := strings.Split(req.URL.Path, "/")[1]
			result, err := showAnimal(res, req, term)
			if err != nil {
				return
			}
			animalList = result

		} else {
			// for root path
			result, err := listAnimals(res, req)
			if err != nil {
				return
			}
			animalList = result
		}
	}

	// Print to screen
	type Result struct {
		Username      string
		ListOfAnimals []Animal
	}
	tpl.ExecuteTemplate(res, "content", Result{username, animalList})
	//	fmt.Fprintln(res, user["name"])
	//	fmt.Fprintln(res, AnimalList[0].Term)

}

func handleUserSession(res http.ResponseWriter, req *http.Request) map[string]string {
	// Getting cookie "sessionid", if nil,
	// creates new sessionid and sets cookie
	cookie, _ := req.Cookie("sessionid")
	if cookie == nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "sessionid",
			Value: id.String(),
		}
		http.SetCookie(res, cookie)
	}

	// New context (only used if serving from app engine)
	ctx := appengine.NewContext(req)

	item, err := memcache.Get(ctx, cookie.Value)
	if err != nil && err != memcache.ErrCacheMiss {
		fmt.Fprintln(res, err)
	}
	// If no error and no item is returned, we create new memcache key value
	if item == nil {

		//Creating JSON data to be store as value
		m := map[string]string{
			"name":  "John Doe",
			"email": "test@example.com",
		}

		// Getting a slice of byte returned from json.Marshal
		// same format needed for cookie value
		bs, _ := json.Marshal(m)

		item = &memcache.Item{
			Key:   cookie.Value,
			Value: bs,
		}

		// Setting key value into memcache
		memcache.Set(ctx, item)
	}

	var m map[string]string
	json.Unmarshal(item.Value, &m)

	return m
}

func handlePost(res http.ResponseWriter, req *http.Request) {
	term := req.FormValue("term")
	definition := req.FormValue("definition")

	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Animal", term, 0, nil)

	entity := Animal{
		Term:       term,
		Definition: definition,
	}

	_, err := datastore.Put(ctx, key, &entity)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return
	}

	http.Redirect(res, req, "/", 302)

}

func showAnimal(res http.ResponseWriter, req *http.Request, term string) ([]Animal, error) {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Animal", term, 0, nil)

	var entity Animal

	err := datastore.Get(ctx, key, &entity)
	if err == datastore.ErrNoSuchEntity {
		tpl.ExecuteTemplate(res, "error", term)
		return nil, err
	} else if err != nil {
		http.Error(res, err.Error(), 500)
		return nil, err
	}
	var list []Animal
	list = append(list, entity)
	return list, nil
}

func listAnimals(res http.ResponseWriter, req *http.Request) ([]Animal, error) {
	ctx := appengine.NewContext(req)
	q := datastore.NewQuery("Animal").Order("Term")

	var list []Animal

	_, err := q.GetAll(ctx, &list)
	if err != nil {
		http.Error(res, err.Error(), 500)
		return nil, err
	}
	return list, nil
}

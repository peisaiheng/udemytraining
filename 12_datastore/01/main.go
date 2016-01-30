package main

import (
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("/", handleIndex)
}

func handleIndex(res http.ResponseWriter, req *http.Request) {
	user := handleUserSession(res, req) // returns map of memcache values
	tpl.ExecuteTemplate(res, "content", user["name"])
	//	fmt.Fprintln(res, user["name"])
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

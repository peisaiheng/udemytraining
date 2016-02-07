package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	// Define routes
	r := httprouter.New()
	r.GET("/", home)
	r.GET("/form/signup", signup)

	// Define APIs
	r.POST("/api/checkusername", checkUsername)
	r.POST("/api/createuser", createUser)
	r.POST("/api/login", loginProcess)
	r.GET("/api/logout", logout)

	// Set router for HTTP
	http.Handle("/", r)
	http.Handle("favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public/"))))

	// parse template
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

// Render templates
func home(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "index.html")
}

func signup(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	serveTemplate(res, req, "signup.html")

}

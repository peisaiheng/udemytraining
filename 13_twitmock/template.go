package main

import (
	"encoding/json"
	"net/http"
)

func serveTemplate(res http.ResponseWriter, req *http.Request, templateName string) {
	memItem, err := getSession(req)
	if err != nil {
		// err, Not logged in
		// execute template with empty session data
		tpl.ExecuteTemplate(res, templateName, SessionData{})
		return
	}

	//logged in
	var sd SessionData
	json.Unmarshal(memItem.Value, &sd)
	sd.LoggedIn = true
	tpl.ExecuteTemplate(res, templateName, &sd)
}
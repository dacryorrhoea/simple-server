package main

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func GetList(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	items, err := getAllItems()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	main := filepath.Join("public", "index.html")

	tmpl, err := template.ParseFiles(main)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = tmpl.ExecuteTemplate(rw, "items", items)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}

package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path != "/about" {
			RenderErrorPage(http.StatusNotFound, w)
			return
		}
		tml, err := template.ParseFiles("templates/about.html")
		if err != nil {
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		tmpl := template.Must(tml, err)
		err = tmpl.Execute(w, nil)
		fmt.Println(err)
		return
	default:
		RenderErrorPage(http.StatusMethodNotAllowed,w)
		return
	}
}

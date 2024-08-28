package handler

import (
	"fmt"
	"forum/backend/dto"
	"forum/backend/server/service"
	"net/http"
	"strings"
	"text/template"
)

func SearchSuggestionHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	q := r.Form.Get("q")
	tokenData, _ := authService.VerifyToken(r)
	if strings.TrimSpace(q) != "" {
		keywords := strings.Fields(q)
		sugg := service.PostSrvice.GetPostByKeywords(keywords, tokenData)
		tml, err := template.ParseFiles("templates/suggestion_layout.html")
		if err != nil {
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		tmpl := template.Must(tml, err)
		err = tmpl.Execute(w, sugg)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		q := r.URL.Query().Get("q")
		if strings.TrimSpace(q) == "" {
			RenderErrorPage(http.StatusBadRequest, w)
			return
		}
		var data Data
		tokenData, err := authService.VerifyToken(r)
		if err == nil {
			data.IsAuthenticated = true
		}
		var posts []dto.PostDTO

		posts = service.PostSrvice.GetPostByKeywords(strings.Fields(q), tokenData)

		data.Posts = posts
		tml, err := template.ParseFiles("templates/search.html")
		if err != nil {
			RenderErrorPage(http.StatusInternalServerError, w)
			return
		}
		tmpl := template.Must(tml, err)
		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}
	}
}

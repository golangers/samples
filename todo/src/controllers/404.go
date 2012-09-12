package controllers

import (
	"net/http"
)

type Page404 struct {
	Application
}

func init() {
	App.SetNotFoundController(Page404{})
}

func (p *Page404) Init(w http.ResponseWriter, r *http.Request) {
	p.Application.Init(w, r)
	p.Document.GenerateHtml = false
	p.Template = "_notfound/404.html"
	w.WriteHeader(http.StatusNotFound)
}

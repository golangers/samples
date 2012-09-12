package controllers

import (
	"golanger.com/web"
	"net/http"
)

type Application struct {
	web.Page
}

func (a *Application) Init(w http.ResponseWriter, r *http.Request) {
	a.Page.Init(w, r)
}

var App = &Application{
	Page: web.NewPage(web.PageParam{}),
}

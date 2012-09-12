package controllers

import (
	"golanger.com/web"
	"net/http"
)

type Application struct {
	web.Page
	RW http.ResponseWriter
	R  *http.Request
}

func (a *Application) Init() {
	a.Page.Init(a.RW, a.R)
}

var App = &Application{
	Page: web.NewPage(web.PageParam{}),
}

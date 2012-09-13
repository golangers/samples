package controllers

import (
	"fmt"
	. "golanger.com/framework/middleware"
	"helper"
	"html/template"
	. "models"
	"net/http"
)

type PageIndex struct {
	Application
}

func init() {
	App.RegisterController("index/", PageIndex{})
}

func (p *PageIndex) Index(w http.ResponseWriter, r *http.Request) {
	mgo := Middleware.Get("db").(*helper.Mongo)
	coll := mgo.C(ColGuestBook)

	query := coll.Find(nil).Sort("-timestamp")

	var entries []ModelGuestBook
	if err := query.All(&entries); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	p.Body = entries
}

func (p *PageIndex) Sign(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		p.Body = "不支持这种请求方式: " + fmt.Sprintf("%v", r.Method)
		p.Template = "index/error.html"
		return
	}

	entry := NewGuestBook()
	entry.Name = p.POST["name"]
	entry.Message = p.POST["message"]

	if entry.Name == "" {
		entry.Name = "Some dummy who forgot a name"
	}
	if entry.Message == "" {
		entry.Message = "Some dummy who forgot a message."
	}

	entry.Name = template.HTMLEscapeString(entry.Name)
	entry.Message = template.HTMLEscapeString(entry.Message)

	mgo := Middleware.Get("db").(*helper.Mongo)
	coll := mgo.C(ColGuestBook)

	if err := coll.Insert(entry); err != nil {
		p.Body = "数据库错误：" + fmt.Sprintf("%v", err)
		p.Template = "index/error.html"
		return
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

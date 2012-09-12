package controllers

import (
	"fmt"
	"net/http"
	"time"
)

type PageIndex struct {
	Application
}

func init() {
	App.RegisterController("index/", PageIndex{})
}

func (p *PageIndex) Index(w http.ResponseWriter, r *http.Request) {
	p.SESSION["String"] = "String"
	p.SESSION["string"] = "string"
	p.SESSION["Int"] = 1
	p.SESSION["Map"] = map[string]string{
		"a": "b",
		"b": "c",
	}
	p.SESSION["Int64"] = time.Now().UnixNano()
}

func (p *PageIndex) TestPage(w http.ResponseWriter, r *http.Request) {
	p.Document.Title = "测试页面"
	w.Write([]byte(fmt.Sprintf("%v", p.SESSION["String"])))
	w.Write([]byte(fmt.Sprintf("%v", p.SESSION["string"])))
	w.Write([]byte(fmt.Sprintf("%v", p.SESSION["Int"])))
	w.Write([]byte(fmt.Sprintf("%v", p.SESSION["Map"])))
	w.Write([]byte(fmt.Sprintf("%v", p.SESSION["Int64"])))
}

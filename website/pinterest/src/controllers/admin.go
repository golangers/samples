package controllers

import (
	"golanger.com/framework/utils"
	. "models"
	"net/http"
	"path/filepath"
	"strconv"
)

type PageAdmin struct {
	Application
}

func init() {
	App.RegisterController("admin/", PageAdmin{})
}

func (p *PageAdmin) Init(w http.ResponseWriter, r *http.Request) {
	p.Application.Init(w, r)
	_, fileName := filepath.Split(r.URL.Path)
	if fileName != "login.html" {
		if _, ok := p.SESSION["user"]; !ok {
			http.Redirect(w, r, "/admin/login.html", http.StatusFound)
		}
	}
}

func (p *PageAdmin) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := p.POST["username"]
		password := p.POST["password"]
		if username == p.Config.Environment["Username"] && password == p.Config.Environment["Password"] {
			p.SESSION["user"] = username
			http.Redirect(w, r, "/admin/index.html", http.StatusFound)
		}
	}
}

func (p *PageAdmin) Logout(w http.ResponseWriter, r *http.Request) {
	delete(p.SESSION, "user")
	http.Redirect(w, r, "/admin/index.html", http.StatusFound)
}

func (p *PageAdmin) Index(w http.ResponseWriter, r *http.Request) {
	body := utils.M{}
	body["invalidImages"], _ = GetInvalidImages()
	body["classes"], _ = GetClasses()

	p.Body = body
}

func (p *PageAdmin) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			var id = p.POST["id"]
			idValue, _ := strconv.ParseInt(id, 0, 64)
			go DeleteImageWithId(idValue)
		}
	}
}

func (p *PageAdmin) Recover(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			var id = p.POST["id"]
			idValue, _ := strconv.Atoi(id)
			image, err := GetImage(idValue)
			if err == nil {
				image.Status = 1
				go SaveImages(*image)
			}

		}
	}
}

func (p *PageAdmin) Adclass(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			var name = p.POST["className"]
			go AddClass(Class{
				Name: name,
			})
		}
	}
}

func (p *PageAdmin) Declass(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			var id = p.POST["classId"]
			idValue, _ := strconv.ParseInt(id, 0, 64)
			go DeleteClassWithId(idValue)
		}
	}
}

func (p *PageAdmin) Edclass(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			var id = p.POST["classId"]
			var name = p.POST["className"]
			idValue, _ := strconv.ParseInt(id, 0, 64)
			go EditClassWithId(idValue, name)
		}
	}
}

package controllers

import (
	. "golanger.com/framework/middleware"
	"golanger.com/framework/utils"
	"golanger.com/framework/web"
	"helper"
	. "models"
	"net/http"
	"net/url"
	"strconv"
)

type Application struct {
	web.Page
	OffLogin bool
	OffRight bool
	RW       http.ResponseWriter
	R        *http.Request
}

var App = &Application{
	Page: web.NewPage(web.PageParam{}),
}

func (a *Application) Init() {
	a.Page.Init(a.RW, a.R)

	if a.OffLogin || a.checkLogin() {
		checkRight, _ := strconv.ParseBool(a.Environment["CheckRight"])

		if checkRight {
			a.getRole()
			a.getModule(false)

			if !a.OffRight {
				a.checkRight()
			}
		} else {
			a.getModule(true)
		}
	}

	a.OffLogin = false
	a.OffRight = false
}

func (a *Application) getRole() {
	if username, nok := a.SESSION[a.Page.Config.M["SESSION_UNAME"].(string)]; nok {
		if _, ok := a.SESSION["role"]; !ok {
			mgoServer := Middleware.Get("db").(*helper.Mongo)
			cols := []ModelRole{}
			colSelector := utils.M{}
			colQuerier := utils.M{"delete": 0, "users": username}
			colSorter := []string{"-right.scope", "-status"}

			query := mgoServer.C(ColRole).Find(colQuerier).Select(colSelector).Sort(colSorter...)
			iter := query.Iter()
			for {
				col := ModelRole{}
				b := iter.Next(&col)
				if b != true {
					break
				}

				cols = append(cols, col)
			}

			a.SESSION["role"] = cols
		}
	}

}

func (a *Application) getModule(showAll bool) {
	if _, ok := a.SESSION["modules"]; !ok {
		mgoServer := Middleware.Get("db").(*helper.Mongo)
		cols := []ModelModule{}
		colSelector := utils.M{"name": 1, "path": 1}
		colQuerier := utils.M{"status": 1}
		colSorter := []string{"-order", "-create_time"}
		hasModule := map[string]bool{}
		roles := a.SESSION["role"]
		query := mgoServer.C(ColModule).Find(colQuerier).Select(colSelector).Sort(colSorter...)
		iter := query.Iter()
		for {
			col := ModelModule{}
			b := iter.Next(&col)
			if b != true {
				break
			}

			if showAll {
				cols = append(cols, col)
			} else {
				for _, role := range roles.([]ModelRole) {
					if _, ok := hasModule[col.Path]; !ok {
						switch role.Right["scope"].(string) {
						case "3": //site
							cols = append(cols, col)
							hasModule[col.Path] = true
						case "2": //app
							cols = append(cols, col)
							hasModule[col.Path] = true
						case "1": //module
							if role.Right["modules"] != nil {
								for _, module := range role.Right["modules"].([]interface{}) {
									mod := module.(utils.M)
									if mod["module"].(string) == col.Path {
										cols = append(cols, col)
										hasModule[col.Path] = true
									}
								}
							}
						case "0": //action
							if role.Right["modules"] != nil {
								for _, module := range role.Right["modules"].([]interface{}) {
									mod := module.(utils.M)
									if len(mod["actions"].([]interface{})) > 0 {
										if mod["module"].(string) == col.Path {
											cols = append(cols, col)
											hasModule[col.Path] = true
										}
									}
								}
							}
						}
					}
				}
			}
		}

		a.SESSION["modules"] = cols
	}
}

func (a *Application) checkRight() {
	var hasRight bool
	reqModule := a.CurrentController
	reqAction := a.CurrentAction
	for _, module := range a.SESSION["modules"].([]ModelModule) {
		if module.Path == reqModule {
			for _, role := range a.SESSION["role"].([]ModelRole) {
				if role.Right["scope"].(string) != "0" {
					hasRight = true
					goto Check_Right
				} else {
					if role.Right["modules"] != nil {
						for _, mod := range role.Right["modules"].([]interface{}) {
							m := mod.(utils.M)
							for _, action := range m["actions"].([]interface{}) {
								if reqAction == action.(string) {
									hasRight = true
									goto Check_Right
								}
							}
						}
					}
				}
			}
		}
	}

Check_Right:
	if !hasRight {
		a.RW.WriteHeader(http.StatusForbidden)
		a.RW.Write([]byte("无权限"))
		a.Close = true
	}
}

func (a *Application) checkLogin() bool {
	var b bool
	if a.checkUser() {
		b = true
		if a.R.URL.Path == "/login.html" {
			http.Redirect(a.RW, a.R, "/index.html", http.StatusFound)
			a.Close = true
		}
	} else {
		if a.R.URL.Path != "/login.html" {
			http.Redirect(a.RW, a.R, "/login.html?back_url="+url.QueryEscape(a.R.URL.String()), http.StatusFound)
			a.Close = true
		}
	}

	return b
}

func (a *Application) checkUser() (res bool) {
	username, uok := a.SESSION[a.Page.Config.M["SESSION_UNAME"].(string)]
	ukey, ukok := a.SESSION[a.Page.Config.M["SESSION_UKEY"].(string)]

	if uok && ukok {
		mgoServer := Middleware.Get("db").(*helper.Mongo)

		colQuerier := utils.M{"name": username, "password": ukey, "status": 1}
		colSelecter := utils.M{"name": 1}
		col := ModelUser{}
		err := mgoServer.C(ColUser).Find(colQuerier).Select(colSelecter).One(&col)

		if err == nil && col.Name != "" {
			res = true
		}
	}

	return res
}

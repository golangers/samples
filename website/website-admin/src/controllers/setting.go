package controllers

import (
	"encoding/json"
	. "golanger.com/framework/middleware"
	"golanger.com/framework/utils"
	"helper"
	. "models"
	"time"
)

type PageSetting struct {
	Application
}

func init() {
	App.RegisterController("setting/", PageSetting{})
}

func (p *PageSetting) Init() {
	p.OffRight = true
	p.Application.Init()
}

func (p *PageSetting) Info() {
	username := p.SESSION[p.M["SESSION_UNAME"].(string)]

	mgoServer := Middleware.Get("db").(*helper.Mongo)
	user := ModelUser{}
	colQuerier := utils.M{
		"delete": 0,
		"name":   username,
	}

	mgoServer.C(ColUser).Find(colQuerier).One(&user)

	p.Body = user
}

func (p *PageSetting) InfoUpdate() {
	username := p.SESSION[p.M["SESSION_UNAME"].(string)]

	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			mgoServer := Middleware.Get("db").(*helper.Mongo)

			m := utils.M{
				"status":  "1",
				"message": "",
			}
			if !ok {
				m["status"] = "0"
				m["message"] = "修改用户时必须指定用户名！"
			}

			colQuerier := utils.M{"name": username, "delete": 0}
			cnt, err := mgoServer.C(ColUser).Find(colQuerier).Count()
			if err != nil || cnt == 0 {
				m["status"] = "0"
				m["message"] = "用户不存在!"
			} else {
				email, emailOk := p.POST["email"]
				password, passwordOk := p.POST["password"]
				updateField, change := utils.M{}, utils.M{}

				if emailOk {
					updateField["email"] = email
				}
				if passwordOk {
					password = utils.Strings(password).Md5()
					updateField["password"] = password
				}
				if emailOk || passwordOk {
					updateField["update_time"] = time.Now().Unix()
					change["$set"] = updateField
					err := mgoServer.C(ColUser).Update(colQuerier, change)
					if err != nil {
						m["status"] = "0"
						m["message"] = "用户资料更新失败！"
					} else {
						m["message"] = "用户资料更新成功！"
					}
				} else {
					m["status"] = "0"
					m["message"] = "请输入需要更新的资料！"
				}
			}
			ret, _ := json.Marshal(m)
			p.RW.Write(ret)

			return
		}
	}
}

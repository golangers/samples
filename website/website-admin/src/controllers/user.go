package controllers

import (
	"encoding/json"
	. "golanger.com/framework/middleware"
	"golanger.com/framework/utils"
	"helper"
	. "models"
	"time"
)

type PageUser struct {
	Application
}

func init() {
	App.RegisterController("user/", PageUser{})
}

// Show user list
func (p *PageUser) Index() {
	mgoServer := Middleware.Get("db").(*helper.Mongo)
	cols := []ModelUser{}
	colSelector := utils.M{}
	colQuerier := utils.M{"delete": 0}
	colSorter := []string{"-create_time"}

	query := mgoServer.C(ColUser).Find(colQuerier).Select(colSelector).Sort(colSorter...)
	iter := query.Iter()
	for {
		col := ModelUser{}
		b := iter.Next(&col)
		if b != true {
			break
		}

		cols = append(cols, col)
	}

	p.Body = cols
}

// Create User
func (p *PageUser) CreateUser() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			mgoServer := Middleware.Get("db").(*helper.Mongo)
			m := utils.M{
				"status":  1,
				"message": "",
			}
			p.Hide = true
			username := p.POST["username"]
			email := p.POST["email"]
			colQuerier := utils.M{"name": username, "delete": 0}
			var col ModelUser
			err := mgoServer.C(ColUser).Find(colQuerier).One(&col)

			if col.Name != "" || err == nil {
				m["status"] = 0
				m["message"] = "该用户名已存在"
			} else {
				password := utils.Strings(p.POST["password"]).Md5()
				tnow := time.Now()
				mgoServer.C(ColUser).Insert(&ModelUser{
					Email:       email,
					Name:        username,
					Password:    password,
					Status:      1,
					Create_time: tnow.Unix(),
					Update_time: tnow.Unix(),
				})
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)

			return
		}
	}
}

// Update User
func (p *PageUser) UpdateUser() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			mgoServer := Middleware.Get("db").(*helper.Mongo)

			username, ok := p.POST["username"]
			m := utils.M{
				"status":  "1",
				"message": "",
			}
			if !ok {
				m["status"] = "0"
				m["message"] = "修改用户时必须指定用户名！"
			}

			colQuerier := utils.M{"name": username}
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

// Delete User
func (p *PageUser) DeleteUser() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			username := p.POST["username"]
			m := utils.M{
				"status":  "1",
				"message": "",
			}

			if username == p.SESSION[p.M["SESSION_UNAME"].(string)].(string) {
				m["status"] = "0"
				m["message"] = "不能删除自己"
			} else {
				mgoServer := Middleware.Get("db").(*helper.Mongo)
				colQuerier := utils.M{"name": username, "delete": 0}
				cnt, err := mgoServer.C(ColUser).Find(colQuerier).Count()
				if err != nil || cnt == 0 {
					m["status"] = "0"
					m["message"] = "用户不存在"
				} else {
					change := utils.M{"$set": utils.M{"delete": 1}}
					err = mgoServer.C(ColUser).Update(colQuerier, change)
					if err != nil {
						m["status"] = "0"
						m["message"] = "删除用户失败！"
					} else {
						m["message"] = "成功删除用户！"
					}
				}
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
		}
	}
}

// Stop User
func (p *PageUser) StopUser() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			m := utils.M{
				"status":  "1",
				"message": "",
			}

			if username, ok := p.POST["username"]; ok {
				if username == p.SESSION[p.M["SESSION_UNAME"].(string)].(string) {
					m["status"] = "0"
					m["message"] = "不能改变自己的状态"
				} else {
					mgoServer := Middleware.Get("db").(*helper.Mongo)
					col := ModelUser{}
					colQuerier := utils.M{"name": username}
					err := mgoServer.C(ColUser).Find(colQuerier).One(&col)
					if err != nil {
						m["status"] = "0"
						m["message"] = "用户不存在"
					} else {
						change := utils.M{}
						if col.Status == 1 {
							change["$set"] = utils.M{"status": 0}
						} else {
							change["$set"] = utils.M{"status": 1}
						}

						err = mgoServer.C(ColUser).Update(colQuerier, change)
						if err != nil {
							m["status"] = "0"
							m["message"] = "停用用户失败！"
						} else {
							m["message"] = "成功停用用户！"
						}
					}
				}
			} else {
				m["status"] = "0"
				m["message"] = "请选择要停用的用户！"
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
		}
	}
}

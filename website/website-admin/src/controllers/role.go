package controllers

import (
	"encoding/json"
	. "golanger.com/middleware"
	"golanger.com/utils"
	"helper"
	. "models"
	"reflect"
	"strings"
	"time"
)

type PageRole struct {
	Application
}

func init() {
	App.RegisterController("role/", PageRole{})
}

func (p *PageRole) Init() {
	p.Application.Init()

	p.Func["strJoin"] = strings.Join
	p.Func["jsonMarshal"] = func(i interface{}) string {
		j, _ := json.Marshal(i)
		return string(j)
	}
}

// Show role list
func (p *PageRole) Index() {
	mgoServer := Middleware.Get("db").(*helper.Mongo)
	cols := []ModelRole{}
	colSelector := utils.M{}
	colQuerier := utils.M{"delete": 0}
	colSorter := []string{"-create_time", "-status"}

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

	colUsers := []ModelUser{}
	colUserSelector := utils.M{}
	colUserQuerier := utils.M{"delete": 0}
	colUserSorter := []string{"-create_time"}

	queryUser := mgoServer.C(ColUser).Find(colUserQuerier).Select(colUserSelector).Sort(colUserSorter...)
	iterUser := queryUser.Iter()
	for {
		col := ModelUser{}
		b := iterUser.Next(&col)
		if b != true {
			break
		}

		colUsers = append(colUsers, col)
	}

	colModules := []ModelModule{}
	colModuleselector := utils.M{}
	colModulesQuerier := utils.M{"status": 1}
	colModulesorter := []string{"-order", "-create_time"}

	queryModules := mgoServer.C(ColModule).Find(colModulesQuerier).Select(colModuleselector).Sort(colModulesorter...)
	iterModules := queryModules.Iter()
	for {
		col := ModelModule{}
		b := iterModules.Next(&col)
		if b != true {
			break
		}

		colModules = append(colModules, col)
	}

	modules := map[string][]string{}
	rAppType := reflect.TypeOf(App)
	for _, module := range colModules {
		if i, ok := p.Controller[module.Path]; ok {
			modules[module.Path] = []string{}
			rvpc := reflect.New(reflect.TypeOf(i))
			rt := rvpc.Type()
			for j := 0; j < rt.NumMethod(); j++ {
				if _, ok := rAppType.MethodByName(rt.Method(j).Name); !ok {
					modules[module.Path] = append(modules[module.Path], rt.Method(j).Name)
				}
			}
		}
	}

	p.Body = utils.M{
		"roles": cols,
		"users": colUsers,
		"join":  strings.Join,
		"getScopeText": func(scope string) string {
			scopes := map[string]string{
				"0": "方法",
				"1": "模块",
				"2": "应用",
				"3": "全站",
			}

			return scopes[scope]
		},
		"models": modules,
	}
}

// Create role
func (p *PageRole) Create() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			mgoServer := Middleware.Get("db").(*helper.Mongo)
			m := utils.M{
				"status":  1,
				"message": "",
			}
			p.Hide = true
			name := p.POST["name"]
			scope := p.POST["scope"]
			colQuerier := utils.M{"name": name, "delete": 0}
			var col ModelRole
			err := mgoServer.C(ColRole).Find(colQuerier).One(&col)

			if col.Name != "" || err == nil {
				m["status"] = 0
				m["message"] = "该角色名已存在"
			} else {
				tnow := time.Now()
				mgoServer.C(ColRole).Insert(&ModelRole{
					Name:   name,
					Status: 1,
					Right: utils.M{
						"scope":   scope,
						"modules": []utils.M{},
					},
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

// Update role
func (p *PageRole) Update() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			mgoServer := Middleware.Get("db").(*helper.Mongo)

			oldName, ook := p.POST["oldname"]
			name, nok := p.POST["name"]
			scope := p.POST["scope"]
			m := utils.M{
				"status":  "1",
				"message": "",
			}
			if !ook {
				m["status"] = "0"
				m["message"] = "请指定角色名！"
			} else if !nok {
				m["status"] = "0"
				m["message"] = "未提交新的角色名！"
			} else {
				colQuerier := utils.M{"name": oldName}
				cnt, err := mgoServer.C(ColRole).Find(colQuerier).Count()
				if err != nil || cnt == 0 {
					m["status"] = "0"
					m["message"] = "角色不存在!"
				} else {
					updateField, change := utils.M{}, utils.M{}
					if name != "" {
						updateField["name"] = name
					}

					updateField["right.scope"] = scope
					updateField["update_time"] = time.Now().Unix()
					change["$set"] = updateField
					err := mgoServer.C(ColRole).Update(colQuerier, change)
					if err != nil {
						m["status"] = "0"
						m["message"] = "角色更新失败！"
					} else {
						m["message"] = "角色更新成功！"
					}
				}
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)

			return
		}
	}
}

// Delete role
func (p *PageRole) Delete() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			name := p.POST["name"]
			m := utils.M{
				"status":  "1",
				"message": "",
			}

			mgoServer := Middleware.Get("db").(*helper.Mongo)
			colQuerier := utils.M{"name": name, "delete": 0}
			cnt, err := mgoServer.C(ColRole).Find(colQuerier).Count()
			if err != nil || cnt == 0 {
				m["status"] = "0"
				m["message"] = "角色不存在"
			} else {
				change := utils.M{"$set": utils.M{"delete": 1}}
				err = mgoServer.C(ColRole).Update(colQuerier, change)
				if err != nil {
					m["status"] = "0"
					m["message"] = "删除角色失败！"
				} else {
					m["message"] = "成功删除角色！"
				}
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
		}
	}
}

// Stop role
func (p *PageRole) Stop() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			m := utils.M{
				"status":  "1",
				"message": "",
			}

			if name, ok := p.POST["name"]; ok {
				mgoServer := Middleware.Get("db").(*helper.Mongo)
				col := ModelRole{}
				colQuerier := utils.M{"name": name}
				err := mgoServer.C(ColRole).Find(colQuerier).One(&col)
				if err != nil {
					m["status"] = "0"
					m["message"] = "角色不存在"
				} else {
					change := utils.M{}
					if col.Status == 1 {
						change["$set"] = utils.M{"status": 0}
					} else {
						change["$set"] = utils.M{"status": 1}
					}

					err = mgoServer.C(ColRole).Update(colQuerier, change)
					if err != nil {
						m["status"] = "0"
						m["message"] = "修改角色状态失败！"
					} else {
						m["message"] = "成功修改角色状态！"
					}
				}
			} else {
				m["status"] = "0"
				m["message"] = "请选择要修改状态的角色！"
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
		}
	}
}

func (p *PageRole) Users() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			m := utils.M{
				"status":  "1",
				"message": "",
			}

			if name, ok := p.POST["name"]; ok {
				mgoServer := Middleware.Get("db").(*helper.Mongo)
				colQuerier := utils.M{"name": name}
				cnt, _ := mgoServer.C(ColRole).Find(colQuerier).Count()
				if cnt == 0 {
					m["status"] = "0"
					m["message"] = "角色不存在"
				} else {
					var users []string
					if v := p.POST["users"]; v != "" {
						users = strings.Split(v, ",")
					}

					change := utils.M{}
					change["$set"] = utils.M{"users": users}

					err := mgoServer.C(ColRole).Update(colQuerier, change)
					if err != nil {
						m["status"] = "0"
						m["message"] = "修改角色状态失败！"
					} else {
						m["message"] = "成功修改角色状态！"
					}
				}
			} else {
				m["status"] = "0"
				m["message"] = "请选择要修改状态的角色！"
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
		}
	}
}

func (p *PageRole) Right() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			m := utils.M{
				"status":  "1",
				"message": "",
			}

			if name, ok := p.POST["name"]; ok {
				mgoServer := Middleware.Get("db").(*helper.Mongo)
				colQuerier := utils.M{"name": name, "delete": 0}
				cnt, _ := mgoServer.C(ColRole).Find(colQuerier).Count()
				if cnt == 0 {
					m["status"] = "0"
					m["message"] = "角色不存在"
				} else {
					var right map[string][]string
					err := json.Unmarshal([]byte(p.POST["right"]), &right)
					modules := []utils.M{}
					if err == nil {
						for k, v := range right {
							modules = append(modules, utils.M{"module": k, "actions": v})
						}

						change := utils.M{
							"$set": utils.M{"right.modules": modules},
						}

						err := mgoServer.C(ColRole).Update(colQuerier, change)
						if err != nil {
							m["status"] = "0"
							m["message"] = "设置角色权限失败！"
						} else {
							m["message"] = "成功设置角色权限！"
						}
					} else {
						m["status"] = "0"
						m["message"] = "设置角色权限失败！"
					}
				}
			} else {
				m["status"] = "0"
				m["message"] = "请选择要设置权限的角色！"
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
		}
	}
}

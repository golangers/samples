package controllers

import (
	"encoding/json"
	. "golanger.com/framework/middleware"
	"golanger.com/framework/utils"
	"helper"
	. "models"
	"time"
)

type PageModule struct {
	Application
}

func init() {
	App.RegisterController("module/", PageModule{})
}

func (p *PageModule) Index() {
	mgoServer := Middleware.Get("db").(*helper.Mongo)
	cols := []ModelModule{}
	colSelector := utils.M{}
	colQuerier := utils.M{}
	colSorter := []string{"-order", "-create_time"}

	query := mgoServer.C(ColModule).Find(colQuerier).Select(colSelector).Sort(colSorter...)
	iter := query.Iter()
	for {
		col := ModelModule{}
		b := iter.Next(&col)
		if b != true {
			break
		}

		cols = append(cols, col)
	}

	p.Body = cols
}

// Create Module
func (p *PageModule) CreateModule() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			m := utils.M{
				"status":  "1",
				"message": "",
			}
			p.Hide = true
			mgoServer := Middleware.Get("db").(*helper.Mongo)
			modulename := p.POST["modulename"]
			modulepath := p.POST["modulepath"]
			colQuerier := utils.M{"path": modulepath, "name": modulename}
			cnt, err := mgoServer.C(ColModule).Find(colQuerier).Count()
			if err != nil {
				m["status"] = 0
				m["message"] = "获取用户信息失败！"
			} else if cnt != 0 {
				m["status"] = 0
				m["message"] = "路径相同的模块已存在"
			} else {
				tnow := time.Now()
				mgoServer.C(ColModule).Insert(&ModelModule{
					Name:        modulename,
					Path:        modulepath,
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

// Delete Module
func (p *PageModule) DeleteModule() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			mgoServer := Middleware.Get("db").(*helper.Mongo)
			modulename := p.POST["modulename"]
			modulepath := p.POST["modulepath"]
			colQuerier := utils.M{"path": modulepath, "name": modulename}
			m := utils.M{
				"status":  "1",
				"message": "",
			}

			err := mgoServer.C(ColModule).Remove(colQuerier)
			if err != nil {
				m["status"] = "0"
				m["message"] = "删除模块失败！"
			} else {
				m["message"] = "成功删除模块！"
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
		}
	}
}

// Stop Module
func (p *PageModule) StopModule() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			m := utils.M{
				"status":  "1",
				"message": "",
			}
			if modulepath, ok := p.POST["modulepath"]; ok {
				mgoServer := Middleware.Get("db").(*helper.Mongo)
				col := ModelModule{}
				colQuerier := utils.M{"path": modulepath}
				err := mgoServer.C(ColModule).Find(colQuerier).One(&col)
				if err != nil {
					m["status"] = "0"
					m["message"] = "模块不存在"
				} else {
					change := utils.M{}
					if col.Status == 1 {
						change["$set"] = utils.M{"status": 0}
					} else {
						change["$set"] = utils.M{"status": 1}
					}

					err = mgoServer.C(ColModule).Update(colQuerier, change)
					if err != nil {
						m["status"] = "0"
						m["message"] = "停用模块失败！"
					} else {
						m["message"] = "成功停用模块！"
					}
				}
			} else {
				m["status"] = "0"
				m["message"] = "请选择要停用的模块！"
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
		}
	}
}

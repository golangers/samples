package controllers

import (
	"encoding/json"
	. "golanger.com/middleware"
	"golanger.com/utils"
	"helper"
	"labix.org/v2/mgo/bson"
	. "models"
	"strconv"
	"time"
)

type PageArticleCategory struct {
	Application
}

func init() {
	App.RegisterController("article_category/", PageArticleCategory{})
}

func (p *PageArticleCategory) Index() {
	body := utils.M{}
	mgoServer := Middleware.Get("db").(*helper.Mongo)
	cols := []ModelArticleCategory{}
	colSelector := utils.M{}
	colQuerier := utils.M{}
	colSorter := []string{"-order", "-parent", "-create_time"}

	query := mgoServer.C(ColArticleCategory).Find(colQuerier).Select(colSelector).Sort(colSorter...)
	iter := query.Iter()
	for {
		col := ModelArticleCategory{}
		b := iter.Next(&col)
		if b != true {
			break
		}

		cols = append(cols, col)
	}

	body["category"] = cols
	body["getCategory"] = func(cid string) string {
		cateText := ""

		for _, v := range cols {
			if cid == v.Id.Hex() {
				cateText = v.Name
				break
			}
		}

		return cateText
	}

	pCates := []ModelArticleCategory{}
	ArticleCategoryQuery(utils.M{"delete": false, "parent": utils.M{"$exists": false}}).All(&pCates)
	body["pCategory"] = pCates

	p.Body = body
}

// Create
func (p *PageArticleCategory) Create() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			m := utils.M{
				"status":  "1",
				"message": "",
			}
			p.Hide = true
			mgoServer := Middleware.Get("db").(*helper.Mongo)
			name := p.POST["name"]
			order, _ := strconv.Atoi(p.POST["order"])
			parentCate := p.POST["parentCate"]
			colQuerier := utils.M{"name": name}
			if parentCate != "" {
				colQuerier["parent"] = bson.ObjectIdHex(parentCate)
			}

			cnt, err := mgoServer.C(ColArticleCategory).Find(colQuerier).Count()
			if err != nil {
				m["status"] = 0
				m["message"] = "获取信息失败！"
			} else if cnt != 0 {
				m["status"] = 0
				m["message"] = "已经存在同名类别"
			} else {
				tnow := time.Now()
				artCate := ModelArticleCategory{
					Name:       name,
					Order:      order,
					Delete:     false,
					CreateTime: tnow.Unix(),
					UpdateTime: tnow.Unix(),
				}
				if parentCate != "" {
					artCate.Parent = bson.ObjectIdHex(parentCate)
				}

				mgoServer.C(ColArticleCategory).Insert(&artCate)
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
			return
		}
	}
}

// Update
func (p *PageArticleCategory) Update() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			m := utils.M{
				"status":  "1",
				"message": "",
			}
			if id, ok := p.POST["id"]; ok {
				mgoServer := Middleware.Get("db").(*helper.Mongo)
				mid := bson.ObjectIdHex(id)
				cnt, _ := mgoServer.C(ColArticleCategory).FindId(mid).Count()
				if cnt == 0 {
					m["status"] = "0"
					m["message"] = "类别不存在"
				} else {
					name := p.POST["name"]
					order, _ := strconv.Atoi(p.POST["order"])
					parentCate := p.POST["parentCate"]
					change := utils.M{}
					change["$set"] = utils.M{"name": name, "order": order, "update_time": time.Now().Unix()}
					if parentCate != "" {
						change["$set"].(utils.M)["parent"] = bson.ObjectIdHex(parentCate)
					}

					err := mgoServer.C(ColArticleCategory).UpdateId(mid, change)
					if err != nil {
						m["status"] = "0"
						m["message"] = "操作失败！"
					} else {
						m["message"] = "操作成功！"
					}
				}
			} else {
				m["status"] = "0"
				m["message"] = "请选择要操作的类别！"
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
		}
	}
}

// Delete
func (p *PageArticleCategory) Delete() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			mgoServer := Middleware.Get("db").(*helper.Mongo)
			id := p.POST["id"]
			m := utils.M{
				"status":  "1",
				"message": "",
			}

			err := mgoServer.C(ColArticleCategory).RemoveId(bson.ObjectIdHex(id))
			if err != nil {
				m["status"] = "0"
				m["message"] = "删除失败！"
			} else {
				m["message"] = "成功删除！"
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
		}
	}
}

// Stop
func (p *PageArticleCategory) Stop() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			p.Hide = true
			m := utils.M{
				"status":  "1",
				"message": "",
			}
			if id, ok := p.POST["id"]; ok {
				mgoServer := Middleware.Get("db").(*helper.Mongo)
				mid := bson.ObjectIdHex(id)
				col := ModelArticleCategory{}
				err := mgoServer.C(ColArticleCategory).FindId(mid).One(&col)
				if err != nil {
					m["status"] = "0"
					m["message"] = "类别不存在"
				} else {
					change := utils.M{}
					if col.Delete {
						change["$set"] = utils.M{"delete": false, "update_time": time.Now().Unix()}
					} else {
						change["$set"] = utils.M{"delete": true, "update_time": time.Now().Unix()}
					}

					err := mgoServer.C(ColArticleCategory).UpdateId(mid, change)
					if err != nil {
						m["status"] = "0"
						m["message"] = "操作失败！"
					} else {
						m["message"] = "操作成功！"
					}
				}
			} else {
				m["status"] = "0"
				m["message"] = "请选择要操作的类别！"
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
		}
	}
}

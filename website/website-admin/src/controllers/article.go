package controllers

import (
	"encoding/json"
	. "golanger.com/middleware"
	"golanger.com/utils"
	"helper"
	"labix.org/v2/mgo/bson"
	. "models"
	"net/http"
	"strconv"
	"time"
)

type PageArticle struct {
	Application
}

func init() {
	App.RegisterController("article/", PageArticle{})
}

func (p *PageArticle) Index() {
	body := utils.M{}
	mgoServer := Middleware.Get("db").(*helper.Mongo)
	cols := []ModelArticle{}
	colSelector := utils.M{}
	colQuerier := utils.M{"delete": false}
	colSorter := []string{"-order", "-create_time"}

	query := mgoServer.C(ColArticle).Find(colQuerier).Select(colSelector).Sort(colSorter...)
	iter := query.Iter()
	for {
		col := ModelArticle{}
		b := iter.Next(&col)
		if b != true {
			break
		}

		cols = append(cols, col)
	}

	catecols := []ModelArticleCategory{}
	catecolSelector := utils.M{}
	catecolQuerier := utils.M{"delete": false}
	catecolSorter := []string{"-order", "parent", "-create_time"}

	catequery := mgoServer.C(ColArticleCategory).Find(catecolQuerier).Select(catecolSelector).Sort(catecolSorter...)
	cateiter := catequery.Iter()
	for {
		catecol := ModelArticleCategory{}
		b := cateiter.Next(&catecol)
		if b != true {
			break
		}

		catecols = append(catecols, catecol)
	}

	body["art"] = cols
	body["category"] = catecols
	body["getCategory"] = func(cid string) string {
		cateText := ""

		for _, v := range catecols {
			if cid == v.Id.Hex() {
				cateText = v.Name
				break
			}
		}

		return cateText
	}

	p.Body = body
}

// Delete
func (p *PageArticle) Delete() {
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
				cnt, _ := mgoServer.C(ColArticle).FindId(mid).Count()
				if cnt == 0 {
					m["status"] = "0"
					m["message"] = "文章不存在"
				} else {
					change := utils.M{}
					change["$set"] = utils.M{"delete": true, "update_time": time.Now().Unix()}

					err := mgoServer.C(ColArticle).UpdateId(mid, change)
					if err != nil {
						m["status"] = "0"
						m["message"] = "操作失败！"
					} else {
						m["message"] = "操作成功！"
					}
				}
			} else {
				m["status"] = "0"
				m["message"] = "请选择要操作的文章！"
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
		}
	}
}

// Stop
func (p *PageArticle) Stop() {
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
				col := ModelArticle{}
				err := mgoServer.C(ColArticle).FindId(mid).One(&col)
				if err != nil {
					m["status"] = "0"
					m["message"] = "文章不存在"
				} else {
					change := utils.M{}
					if col.Status == 1 {
						change["$set"] = utils.M{"status": 0, "update_time": time.Now().Unix()}
					} else {
						change["$set"] = utils.M{"status": 1, "update_time": time.Now().Unix()}
					}

					err := mgoServer.C(ColArticle).UpdateId(mid, change)
					if err != nil {
						m["status"] = "0"
						m["message"] = "操作失败！"
					} else {
						m["message"] = "操作成功！"
					}
				}
			} else {
				m["status"] = "0"
				m["message"] = "请选择要操作的文章！"
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
		}
	}
}

// Create
func (p *PageArticle) Create() {
	if p.R.Method == "POST" {
		if _, ok := p.POST["ajax"]; ok {
			m := utils.M{
				"status":  "1",
				"message": "",
			}
			p.Hide = true
			mgoServer := Middleware.Get("db").(*helper.Mongo)
			title := p.POST["title"]
			top, _ := strconv.Atoi(p.POST["top"])
			hot, _ := strconv.Atoi(p.POST["hot"])
			order, _ := strconv.Atoi(p.POST["order"])
			target_href := p.POST["target_href"]
			category := p.POST["category"]
			summary := p.POST["summary"]
			content := p.POST["content"]
			colQuerier := utils.M{"category": category, "title": title}
			cnt, err := mgoServer.C(ColArticle).Find(colQuerier).Count()
			if err != nil {
				m["status"] = 0
				m["message"] = "获取信息失败！"
			} else if cnt != 0 {
				m["status"] = 0
				m["message"] = "在同一类别下已经存在同名文章"
			} else {
				tnow := time.Now()
				mgoServer.C(ColArticle).Insert(&ModelArticle{
					Title:      title,
					Top:        top,
					Hot:        hot,
					Order:      order,
					TargetHref: target_href,
					Category:   bson.ObjectIdHex(category),
					Summary:    summary,
					Content:    content,
					Status:     1,
					CreateTime: tnow.Unix(),
					UpdateTime: tnow.Unix(),
				})
			}

			ret, _ := json.Marshal(m)
			p.RW.Write(ret)
			return
		}
	}
}

// Update
func (p *PageArticle) Update() {
	if p.R.Method == "POST" {
		if id, ok := p.POST["id"]; ok {
			mid := bson.ObjectIdHex(id)
			col := ModelArticle{}
			mgoServer := Middleware.Get("db").(*helper.Mongo)
			err := mgoServer.C(ColArticle).FindId(mid).One(&col)
			if err != nil {
				p.RW.Write([]byte("无此文章"))
			} else {
				title := p.POST["title"]
				if title == "" {
					p.SESSION["__ONCE"] = map[string]string{
						"type":    "error",
						"message": "文章标题不得为空",
					}
				} else {
					top, _ := strconv.Atoi(p.POST["top"])
					hot, _ := strconv.Atoi(p.POST["hot"])
					order, _ := strconv.Atoi(p.POST["order"])
					target_href := p.POST["target_href"]
					category := p.POST["category"]
					summary := p.POST["summary"]
					content := p.POST["content"]
					change := utils.M{}
					change["$set"] = utils.M{
						"title":       title,
						"top":         top,
						"hot":         hot,
						"order":       order,
						"target_href": target_href,
						"category":    category,
						"summary":     summary,
						"content":     content,
						"update_time": time.Now().Unix(),
					}

					err := mgoServer.C(ColArticle).UpdateId(mid, change)
					if err == nil {
						p.SESSION["__ONCE"] = map[string]string{
							"type":    "success",
							"message": "修改文章成功",
						}
					}
				}

				p.Hide = true
				http.Redirect(p.RW, p.R, p.R.Referer(), http.StatusFound)

				return
			}
		}
	}

	if id, ok := p.GET["id"]; ok {
		body := utils.M{}
		catecols := []ModelArticleCategory{}
		catecolSelector := utils.M{}
		catecolQuerier := utils.M{"delete": false}
		catecolSorter := []string{"-order", "-create_time"}
		mgoServer := Middleware.Get("db").(*helper.Mongo)
		catequery := mgoServer.C(ColArticleCategory).Find(catecolQuerier).Select(catecolSelector).Sort(catecolSorter...)
		cateiter := catequery.Iter()
		for {
			catecol := ModelArticleCategory{}
			b := cateiter.Next(&catecol)
			if b != true {
				break
			}

			catecols = append(catecols, catecol)
		}

		body["category"] = catecols

		mid := bson.ObjectIdHex(id)
		artContent := ModelArticle{}
		err := mgoServer.C(ColArticle).FindId(mid).One(&artContent)
		if err == nil {
			body["art"] = artContent
		}

		p.Title = "修改文章:" + artContent.Title

		p.Body = body
	}
}

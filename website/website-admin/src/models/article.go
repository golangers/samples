package models

import (
	"errors"
	. "golanger.com/middleware"
	"golanger.com/utils"
	"helper"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

var (
	ColArticle = utils.M{
		"name": "article",
		"index": []string{
			"category,status,delete,-order,-create_time",
			"category,delete,-top,-hot,-create_time",
		},
		"unique": []string{
			"category,title",
		},
	}
)

/*
文章表
article {
    "id" : <id>,
    "category" : <category_id>,
    "title" : <title>,
    "author" : <author>,
    "target_href" : <target_href>,
    "summary" : <summary>,
    "content" : <content>,
    "top" : <top>,
    "hot" : <hot>,
    "order" : <order>,
    "status" : <status>,
    "delete" : <delete>,
    "create_time" : <create_time>,
    "update_time" : <update_time>,
}
*/
type ModelArticle struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	Category   bson.ObjectId `bson:"category"`
	Title      string        `bson:"title"`
	Author     string        `bson:"author"`
	TargetHref string        `bson:"target_href"`
	Summary    string        `bson:"summary"`
	Content    string        `bson:"content"`
	Top        int           `bson:"top"`
	Hot        int           `bson:"hot"`
	Order      int           `bson:"order"`
	Status     int           `bson:"status"`
	Delete     bool          `bson:"delete"`
	CreateTime int64         `bson:"create_time"`
	UpdateTime int64         `bson:"update_time"`
}

func ArticleQuery(querier utils.M) *mgo.Query {
	mgoServer := Middleware.Get("db").(*helper.Mongo)
	return mgoServer.C(ColArticle).Find(querier)
}

func ArticleById(id string, bs ...bool) (ModelArticle, error) {
	col := ModelArticle{}
	var err error
	lbs := len(bs)
	if lbs == 0 {
		err = ArticleQuery(utils.M{"_id": bson.ObjectIdHex(id)}).One(&col)
	} else {
		m := utils.M{"_id": bson.ObjectIdHex(id)}
		if lbs > 0 {
			m["delete"] = bs[0]
		}

		err = ArticleQuery(m).One(&col)
	}

	return col, err
}

func (m *ModelArticle) Add() error {
	if m.Id != "" {
		return errors.New("Had Id")
	}

	//manual control set _id value
	m.Id = bson.NewObjectId()

	m.CreateTime = time.Now().Unix()
	m.UpdateTime = m.CreateTime
	m.Content = helper.FilterHtmlTag("script|noscript|form|fieldset|iframe|math|ins|del", m.Content)
	mgoServer := Middleware.Get("db").(*helper.Mongo)
	err := mgoServer.C(ColArticle).Insert(m)
	if err == nil {
		obj := ModelArticle{}
		obj, err = ArticleById(m.Id.Hex())
		if err == nil {
			*m = obj
		}
	}

	return err
}

func (m *ModelArticle) Remove() error {
	mgoServer := Middleware.Get("db").(*helper.Mongo)
	err := mgoServer.C(ColArticle).RemoveId(m.Id)
	if err == nil {
		m = &ModelArticle{}
	}

	return err
}

func (m *ModelArticle) Update(obj ModelArticle) error {
	change := utils.M{}
	obj.Content = helper.FilterHtmlTag("script|noscript|form|fieldset|iframe|math|ins|del", obj.Content)

	objMap := utils.Struct{obj}.StructToSnakeKeyMap()
	delete(objMap, "id")
	delete(objMap, "create_time")
	objMap["update_time"] = time.Now().Unix()
	change["$set"] = objMap
	mgoServer := Middleware.Get("db").(*helper.Mongo)
	err := mgoServer.C(ColArticle).UpdateId(m.Id, change)
	if err == nil {
		nm, e := ArticleById(m.Id.Hex())
		err = e
		if err == nil {
			*m = nm
		}
	}

	return err
}

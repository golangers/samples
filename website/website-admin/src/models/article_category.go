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
	ColArticleCategory = utils.M{
		"name": "article_category",
		"index": []string{
			"name",
			"parent,delete,-order,-create_time",
		},
		"unique": []string{
			"parent,name",
		},
	}
)

/*
文章分类表
article_category {
    "id" : <id>,
    "parent" : <parent_id>,
    "name" : <name>,
    "order" : <order>,
    "delete" : <delete>,
    "create_time" : <create_time>,
    "update_time" : <update_time>,
}
*/
type ModelArticleCategory struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	Parent     bson.ObjectId `bson:"parent,omitempty"`
	Name       string        `bson:"name"`
	Order      int           `bson:"order"`
	Delete     bool          `bson:"delete"`
	CreateTime int64         `bson:"create_time"`
	UpdateTime int64         `bson:"update_time"`
}

func ArticleCategoryQuery(querier utils.M) *mgo.Query {
	mgoServer := Middleware.Get("db").(*helper.Mongo)
	return mgoServer.C(ColArticleCategory).Find(querier)
}

func ArticleCategoryById(id string, bs ...bool) (ModelArticleCategory, error) {
	col := ModelArticleCategory{}
	var err error
	lbs := len(bs)
	if lbs == 0 {
		err = ArticleCategoryQuery(utils.M{"_id": bson.ObjectIdHex(id)}).One(&col)
	} else {
		m := utils.M{"_id": bson.ObjectIdHex(id)}
		if lbs > 0 {
			m["delete"] = bs[0]
		}

		err = ArticleCategoryQuery(m).One(&col)
	}

	return col, err
}

func (m *ModelArticleCategory) Add() error {
	if m.Id != "" {
		return errors.New("Had Id")
	}

	//manual control set _id value
	m.Id = bson.NewObjectId()

	m.CreateTime = time.Now().Unix()
	m.UpdateTime = m.CreateTime
	mgoServer := Middleware.Get("db").(*helper.Mongo)
	err := mgoServer.C(ColArticleCategory).Insert(m)
	if err == nil {
		obj := ModelArticleCategory{}
		obj, err = ArticleCategoryById(m.Id.Hex())
		if err == nil {
			*m = obj
		}
	}

	return err
}

func (m *ModelArticleCategory) Remove() error {
	mgoServer := Middleware.Get("db").(*helper.Mongo)
	err := mgoServer.C(ColArticleCategory).RemoveId(m.Id)
	if err == nil {
		m = &ModelArticleCategory{}
	}

	return err
}

func (m *ModelArticleCategory) Update(obj ModelArticleCategory) error {
	change := utils.M{}

	objMap := utils.Struct{obj}.StructToSnakeKeyMap()
	delete(objMap, "id")
	delete(objMap, "create_time")
	objMap["update_time"] = time.Now().Unix()
	change["$set"] = objMap
	mgoServer := Middleware.Get("db").(*helper.Mongo)
	err := mgoServer.C(ColArticleCategory).UpdateId(m.Id, change)
	if err == nil {
		nm, e := ArticleCategoryById(m.Id.Hex())
		err = e
		if err == nil {
			*m = nm
		}
	}

	return err
}

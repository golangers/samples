package models

import (
	"encoding/gob"
	"golanger.com/utils"
)

func init() {
	gob.Register(ModelModule{})
	gob.Register([]ModelModule{})
}

var (
	ColModule = utils.M{
		"name": "module",
		"index": []string{
			"path,name,order,status",
		},
	}
)

/*
模块表
module {
    "name" : <name>,
    "path" : <path>,
    "order" : <order>,
    "status" : <status>,
    "create_time" : <create_time>,
    "update_time" : <update_time>,
}
*/
type ModelModule struct {
	Name        string `bson:"name"`
	Path        string `bson:"path"`
	Order       int    `bson:"order"`
	Status      byte   `bson:"status"`
	Create_time int64  `bson:"create_time"`
	Update_time int64  `bson:"update_time"`
}

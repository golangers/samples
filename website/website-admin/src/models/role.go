package models

import (
	"encoding/gob"
	"golanger.com/utils"
)

func init() {
	gob.Register(ModelRole{})
	gob.Register([]ModelRole{})
}

var (
	ColRole = utils.M{
		"name": "role",
		"index": []string{
			"name,status,right.scope,right.modules.module,delete",
		},
	}
)

/*
角色表
role {
    "name" : <name>,
    "users" : []<user_name>,
    "status" : <status>,
    "right" : {
        "scope" : <scope>,// 0=>action,1=>module,2=>app,3=>site
        "modules" : []{
            "module":<module_path>,
            "actions" : []<action_name>,
        }
    },
    "delete" : <delete>,
    "create_time" : <create_time>,
    "update_time" : <update_time>,
}
*/
type ModelRole struct {
	Name        string   `bson:"name"`
	Users       []string `bson:"users"`
	Status      byte     `bson:"status"`
	Right       utils.M  `bson:"right"`
	Delete      byte     `bson:"delete"`
	Create_time int64    `bson:"create_time"`
	Update_time int64    `bson:"update_time"`
}

package models

import (
	"golanger.com/framework/utils"
)

var (
	ColRole = utils.M{
		"name":  "role",
		"index": []string{"name", "status", "right.scope", "right.modules.module", "delete"},
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
	Name        string
	Users       []string
	Status      byte
	Right       utils.M
	Delete      byte
	Create_time int64
	Update_time int64
}

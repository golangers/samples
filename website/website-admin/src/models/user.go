package models

import (
	"golanger.com/framework/utils"
)

var (
	ColUser = utils.M{
		"name":  "user",
		"index": []string{"name", "email", "status", "delete"},
	}
)

/*
用户表
user {
    "name" : <name>,
    "email" : <email>,
    "password" : <password>,
    "status" : <status>,
    "delete" : <delete>,
    "create_time" : <create_time>,
    "update_time" : <update_time>,
}
*/
type ModelUser struct {
	Name        string
	Email       string
	Password    string
	Status      byte
	Delete      byte
	Create_time int64
	Update_time int64
}

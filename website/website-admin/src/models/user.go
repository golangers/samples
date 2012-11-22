package models

import (
	"encoding/gob"
	"golanger.com/framework/utils"
)

func init() {
	gob.Register(ModelUser{})
	gob.Register([]ModelUser{})
}

var (
	ColUser = utils.M{
		"name": "user",
		"index": []string{
			"name,email,status,delete",
		},
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
	Name        string `bson:"name"`
	Email       string `bson:"email"`
	Password    string `bson:"password"`
	Status      byte   `bson:"status"`
	Delete      byte   `bson:"delete"`
	Create_time int64  `bson:"create_time"`
	Update_time int64  `bson:"update_time"`
}

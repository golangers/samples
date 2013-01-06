package models

import (
	"golanger.com/utils"
	"labix.org/v2/mgo/bson"
	"time"
)

var (
	ColGuestBook = utils.M{
		"name":  "guest_book",
		"index": []string{"_id", "name"},
	}
)

type ModelGuestBook struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Timestamp time.Time
	Name      string
	Message   string
}

func NewGuestBook() *ModelGuestBook {
	return &ModelGuestBook{
		Timestamp: time.Now(),
	}
}

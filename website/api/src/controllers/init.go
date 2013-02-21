package controllers

import (
	"golanger.com/webrouter"
)

func init() {
	webrouter.Register("/", &PageIndex{})
}

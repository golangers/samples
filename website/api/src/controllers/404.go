package controllers

import (
	"golanger.com/webrouter"
)

func init() {
	webrouter.NotFoundHtmlHandler(`404`)
}

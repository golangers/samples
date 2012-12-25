package controllers

import (
	"golanger.com/framework/web"
	"net/http"
)

type Application struct {
	web.Page
	//带有这两个成员变量，框架会自动隐示的赋值
	//可以不带这两个成员变量。
	//如果不带，Init方法需要显示传入http.ResponseWriter和*http.Request
	//注意: http.ResponseWriter和*http.Request同时存在时，参数http.ResponseWriter必须在参数*http.Request前面
	//如: Init(rw http.ResponseWriter, req *http.Request)
	//并且每一个控制器如果需要用到http.ResponseWriter或者*http.Request时需要显示的带参数
	//RW 变量名固定
	RW http.ResponseWriter
	//R 变量名固定
	R *http.Request
}

func (a *Application) Init() {
	//Page.Init无论如何，都需要显示的传入http.ResponseWriter和*http.Request
	a.Page.Init(a.RW, a.R)
}

var App = &Application{
	Page: web.NewPage(web.PageParam{}),
}

package controllers

import (
	. "golanger.com/framework/middleware"
	"golanger.com/framework/utils"
	"helper"
	"io/ioutil"
	. "models"
	"net/http"
	"os"
	"time"
)

type PageInstall struct {
	Application
}

func init() {
	App.RegisterController("install/", PageInstall{})
}

func (p *PageInstall) Init() {
	p.OffLogin = true
	p.OffRight = true
	p.Application.Init()
}

func (p *PageInstall) Index() {
	fileInstallLock := "./data/install.lock"

	if _, err := os.Stat(fileInstallLock); err == nil {
		p.RW.Write([]byte("程序已经安装过，如需要重新安装，请删除data目录下的install.lock文件后重试"))
	} else {
		mgoServer := Middleware.Get("db").(*helper.Mongo)
		email := "download@golanger.com"
		username := "golanger"
		password := "leetaifook"
		passwordMd5 := utils.Strings(password).Md5()
		tnow := time.Now()
		mgoServer.C(ColUser).Insert(&ModelUser{
			Email:       email,
			Name:        username,
			Password:    passwordMd5,
			Status:      1,
			Create_time: tnow.Unix(),
			Update_time: tnow.Unix(),
		})

		mgoServer.C(ColModule).Insert(&ModelModule{
			Name:        "模块管理",
			Path:        "module/",
			Order:       0,
			Status:      1,
			Create_time: tnow.Unix(),
			Update_time: tnow.Unix(),
		})

		mgoServer.C(ColModule).Insert(&ModelModule{
			Name:        "用户管理",
			Path:        "user/",
			Order:       0,
			Status:      1,
			Create_time: tnow.Unix(),
			Update_time: tnow.Unix(),
		})

		mgoServer.C(ColModule).Insert(&ModelModule{
			Name:        "角色管理",
			Path:        "role/",
			Order:       0,
			Status:      1,
			Create_time: tnow.Unix(),
			Update_time: tnow.Unix(),
		})

		mgoServer.C(ColRole).Insert(&ModelRole{
			Name:   "超级管理员",
			Users:  []string{username},
			Status: 1,
			Right: utils.M{
				"scope":   "3",
				"modules": []utils.M{},
			},
			Create_time: tnow.Unix(),
			Update_time: tnow.Unix(),
		})

		ioutil.WriteFile(fileInstallLock, []byte("installed"), 0777)

		sessionSign := p.COOKIE[p.Session.CookieName]
		if sessionSign != "" {
			p.Session.Clear(sessionSign)
		}

		p.RW.Write([]byte("安装成功...<br/>用户名:" + username + ",密码:" + password + "<br/>"))
	}

	http.Redirect(p.RW, p.R, "/login.html", http.StatusFound)

	p.Close = true
}

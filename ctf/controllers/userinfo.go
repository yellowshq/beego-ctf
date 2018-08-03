package controllers

import (
	"github.com/astaxie/beego"
)

type UserInfoController struct {
	beego.Controller
}

func (c *UserInfoController) Get(){
	c.TplName="page/UserInfo.html"
}
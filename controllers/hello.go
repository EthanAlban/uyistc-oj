package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	c.TplName = "login.tpl"
}

// Post 接受post请求
func (c *HelloController) Post() {
	c.TplName = "login.tpl"
}

package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {
	c.TplName = "goods.tpl"
}

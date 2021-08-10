package routers

import "unioj/controllers"

import (
	"github.com/astaxie/beego"
)

func annoncementRouter() {
	beego.Router("api/annonce/get_annoncements", &controllers.AnnoncementController{}, "post:GetAnnoncements")
}

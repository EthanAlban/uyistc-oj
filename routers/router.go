package routers

import (
	"beegoDemo01/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/goods", &controllers.GoodsController{})
	beego.Router("/hello", &controllers.HelloController{})
}

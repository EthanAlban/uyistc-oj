package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers"
)

func contestsRouter() {
	beego.Router("api/contests/get_all_contests", &controllers.ContestsController{}, "get:GetAllContests")
	beego.Router("api/contests/get_contest_quantity", &controllers.ContestsController{}, "get:GetContestQuantity")
}

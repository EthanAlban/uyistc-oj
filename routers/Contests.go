package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers"
)

func contestsRouter() {
	beego.Router("api/contests/get_all_contests", &controllers.ContestsController{}, "get:GetAllContests")
	beego.Router("api/contests/get_contest_quantity", &controllers.ContestsController{}, "get:GetContestQuantity")
	beego.Router("api/contests/contest_access", &controllers.ContestsController{}, "get:ContestAccess")
	beego.Router("api/contests/contest_password_check", &controllers.ContestsController{}, "get:ContestPasswordCheck")
	beego.Router("api/contests/contest_problem_list", &controllers.ContestsController{}, "get:ContestProblemList")
}

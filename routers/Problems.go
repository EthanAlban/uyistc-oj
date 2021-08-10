package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers"
)

func problemsRouter() {
	beego.Router("api/problems/get_pages_problems", &controllers.ProblemsController{}, "post:GetPagesProblems")
}

package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers"
)

func problemsRouter() {
	beego.Router("api/problems/get_pages_problems", &controllers.ProblemsController{}, "post:GetPagesProblems")
	beego.Router("api/problems/get_problem_detail_by_id", &controllers.ProblemsController{}, "post:GetProblemDetailById")
	beego.Router("api/problems/get_problems_by_tag", &controllers.ProblemsController{}, "get:GetProblemByTag")
}

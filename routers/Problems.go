package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers"
)

func problemsRouter() {
	beego.Router("api/problems/get_pages_problems", &controllers.ProblemsController{}, "get:GetPagesProblems")
	beego.Router("api/problems/get_problem_detail_by_id", &controllers.ProblemsController{}, "get:GetProblemDetailById")
	beego.Router("api/problems/get_problems_by_tag", &controllers.ProblemsController{}, "get:GetProblemByTag")
	beego.Router("api/problems/get_problem_tags_by_id", &controllers.ProblemsController{}, "get:GetProblemTagsById")
	beego.Router("api/problems/get_problem_ac_sub_times", &controllers.ProblemsController{}, "get:GetProblemAcSubTimes")
	beego.Router("api/problems/get_problem_support_language", &controllers.ProblemsController{}, "get:GetProblemSupportLanguage")
}

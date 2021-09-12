package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers"
)

func judgerRouter() {
	beego.Router("api/judger/new_judge_task", &controllers.SubmissionController{}, "post:SendTaskToKafka")
	beego.Router("api/submission/submision_exist", &controllers.SubmissionController{}, "get:IsSubmissionExsit")
	beego.Router("api/submission/get_submission_final", &controllers.SubmissionController{}, "get:GetFinalInfoOfSubmission")
	beego.Router("api/submission/get_submission_static_for_problem", &controllers.SubmissionController{}, "get:GetSubmissionStaticForProblem")
}

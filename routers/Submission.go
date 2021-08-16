package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers"
)

func judgerRouter() {
	beego.Router("api/judger/new_judge_task", &controllers.SubmissionController{}, "post:SendTaskToKafka")
}

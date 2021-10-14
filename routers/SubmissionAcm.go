package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers"
)

func submissionAcmRouter() {
	beego.Router("api/acm/send_task_to_kafka", &controllers.SubmissionAcmController{}, "get:SendTaskToKafka")
}

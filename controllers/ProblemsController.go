package controllers

import (
	"fmt"
	"unioj/models"
)

type ProblemsController struct {
	BaseController
}

func (this *ProblemsController) GetPagesProblems() {
	pros, err := models.NewProblems().GetPagesProblems(10, 0)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(*pros); i++ {
		models.O.LoadRelated(&((*pros)[i]), "ProblemType")
		(*pros)[i].Uid, _ = models.NewUser().GetUserByUid((*pros)[i].Uid.UId)
	}
	this.JsonResult(200, "问题列表加载成功", pros)
}

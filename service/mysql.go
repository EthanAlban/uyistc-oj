package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wonderivan/logger"
)

// Conn 用于在得到判题结果之后更新submission的内容
var Conn *sql.DB

func InitMysqlServer(user, password, host string) (err error) {
	Conn, err = sql.Open("mysql", user+":"+password+"@tcp("+host+")/unioj?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		logger.Error("[3] ERROR:初始化连接mysql失败...  error:", err)
		//fmt.Println("[3] ERROR:初始化连接mysql失败...  error:", err)
		return err
	}
	logger.Debug("[3] INFO:初始化连接mysql成功... ")
	//fmt.Println("[3] INFO:初始化连接mysql成功... ")
	return err
}

func UpdateSubmission(submision Submission) {
	ret, err := Conn.Exec("UPDATE submission set result=?,err_info=? where submission_id=?", submision.Result, submision.ErrInfo, submision.SubmissionId)
	if err != nil {
		logger.Error("更新问题失败,err:", err)
	}
	updNums, _ := ret.RowsAffected()
	logger.Info("更新记录数量:", updNums, " 更新问题:", submision.ProblemId.Title)
	//fmt.Println("更新记录数量:", updNums)
}

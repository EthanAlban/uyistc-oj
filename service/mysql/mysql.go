package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wonderivan/logger"
	"unioj-Judger/service/structs"
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

func UpdateSubmission(submision structs.Submission) {
	ret, err := Conn.Exec("UPDATE submission set result=?,err_info=? where submission_id=?", submision.Result, submision.ErrInfo, submision.SubmissionId)
	if err != nil {
		logger.Error("更新问题失败,err:", err)
	}
	updNums, _ := ret.RowsAffected()
	logger.Info("更新记录数量:", updNums, " 更新问题:", submision.ProblemId.Title)
	//fmt.Println("更新记录数量:", updNums)
}

// CheckJudgerExsit 检查对应的ip:port有没有已经注册了的判题器
func CheckJudgerExsit(ip_port string) int {
	var nums int64
	err := Conn.QueryRow("SELECT  count(*) FROM judger where judger_host=?", ip_port).Scan(&nums)
	if err != nil {
		logger.Error("查询已有判题器失败：", err)
	}
	logger.Info("查询已有判题器成功...  nums：", nums)
	return int(nums)
}

func RegisterJudgerIntoSql(ip_port, topic string) {
	_, err := Conn.Exec("insert into judger(judger_host,online,kafka_topic) values(?,?,?)", ip_port, 1, topic)
	if err != nil {
		logger.Error("注册判题器失败：", err)
	}
	logger.Info("注册判题器成功...")
}

// UpdateJudger 更新当前存在的判题器状态和topic
func UpdateJudger(ip_port, topics string) {
	_, err := Conn.Exec("UPDATE judger set online=?,kafka_topic=? where judger_host=?", 1, topics, ip_port)
	if err != nil {
		logger.Error("更新判题器 ", ip_port, " 状态失败...  err:", err)
	}
	logger.Info("更新判题器 ", ip_port, " 状态成功...")
}

// DestroyJudger 清除判题器记录
func DestroyJudger(ip_port string) {
	fmt.Println("------------------------------------")
	_, err := Conn.Exec("delete from judger where judger_host=?", ip_port)
	if err != nil {
		logger.Error("销毁判题器 ", ip_port, " 状态失败...  err:", err)
	}
	logger.Info("销毁判题器 ", ip_port, " 状态成功...")
}

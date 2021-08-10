package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	logger "unioj/logs"
	"unioj/models"
	"unioj/utils"
)

type UserController struct {
	BaseController
}

type pramas struct {
	UserName string
	Password string
	Email    string
}

func (this *UserController) GetuserbyUsernamePwd() {
	var page pramas
	this.ParseForm(&page)
	var user, err = models.NewUser().GetuserbyUsernamePwd(page.UserName, page.Password, "email")
	if user == nil {
		this.JsonResult(404, "can't find such user", err)
	}
	this.JsonResult(0, "ok", user)
}

// Login 用户登录api
func (this *UserController) Login() {
	var params pramas
	body := this.Ctx.Input.RequestBody //这是获取到的json二进制数据
	json.Unmarshal(body, &params)      //解析二进制json，把结果放进ob中
	//判断用户名类型
	var usernameType string
	if utils.VerifyEmailFormat(params.UserName) {
		usernameType = "email"
	} else if utils.VerifyMobileFormat(params.UserName) {
		usernameType = "tel"
	} else {
		usernameType = "username"
	}
	user, err := models.NewUser().GetuserbyUsernamePwd(params.UserName, params.Password, usernameType)
	if err != nil {
		fmt.Printf("%v ", err.Error())
		if err.Error() == "<QuerySeter> no row found" {
			this.JsonResult(204, "没有找到相关用户", err.Error())
		} else if err.Error() == "wrong password" {
			this.JsonResult(205, "用户名密码不匹配", err.Error())
		}
	}
	//记录session
	if this.Ctx.Input.CruSession == nil {
		this.StartSession()
	}
	this.Ctx.Input.CruSession.Set("userid", user.UId)
	//sess,err := session.GlobalSessions.SessionStart(this.Ctx.ResponseWriter,this.Ctx.Request)
	//err = sess.Set("userid", user.UId)
	if err != nil {
		fmt.Println(" err:" + err.Error())
	}
	this.JsonResult(200, "登陆成功", user)
}

// Register 新用户注册
func (this *UserController) Register() {
	var params pramas
	body := this.Ctx.Input.RequestBody
	err := json.Unmarshal(body, &params)
	if err != nil {
		fmt.Println(err)
	}
	var user models.User
	user.UserName = params.UserName
	user.Email = params.Email
	user.Password = params.Password
	if models.NewUser().IsUserNameUsed(params.UserName) {
		this.JsonResult(205, "username already registered")
	}
	if models.NewUser().IsEmailUsed(params.Email) {
		this.JsonResult(206, "email already registered")
	}
	_, err = models.O.Insert(&user)
	if err != nil {
		fmt.Println("注册用户：" + params.UserName + "失败，err:" + err.Error())
		logger.LogError("注册用户：" + params.UserName + "失败，err:" + err.Error())
	}
	this.JsonResult(200, "登陆成功", user)
}

// UserProfile 按照session中存储的userid查询登陆用户的信息
func (this *UserController) UserProfile() {
	userid := this.Ctx.Input.CruSession.Get("userid")
	var user *models.User
	if userid != nil {
		user, _ = models.NewUser().GetUserByUid(userid.(int))
	}
	this.JsonResult(200, "msg", user)
}

// Logout 用户登出（退出登陆）
func (this *UserController) Logout() {
	this.Ctx.Input.CruSession.Flush()
	this.Ctx.Input.CruSession.Set("userid", nil)
	this.Ctx.Input.CruSession = nil
	beego.GlobalSessions.SessionDestroy(this.Ctx.ResponseWriter, this.Ctx.Request)
	this.JsonResult(200, "登出")
}

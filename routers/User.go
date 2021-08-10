package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers"
)

func userRouter() {
	beego.Router("api/usr/getuser_by_username_pwd", &controllers.UserController{}, "get:GetuserbyUsernamePwd")
	beego.Router("api/usr/login", &controllers.UserController{}, "post:Login")
	beego.Router("api/usr/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("api/usr/register", &controllers.UserController{}, "post:Register")
	beego.Router("api/usr/user_profile", &controllers.UserController{}, "get:UserProfile")
}

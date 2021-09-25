package routers

import (
	"github.com/astaxie/beego"
	"unioj/controllers/users"
)

func userRouter() {
	beego.Router("api/usr/getuser_by_username_pwd", &users.UserController{}, "get:GetuserbyUsernamePwd")
	beego.Router("api/usr/login", &users.UserController{}, "post:Login")
	beego.Router("api/usr/logout", &users.UserController{}, "get:Logout")
	beego.Router("api/usr/register", &users.UserController{}, "post:Register")
	beego.Router("api/usr/user_profile", &users.UserController{}, "get:UserProfile")
}

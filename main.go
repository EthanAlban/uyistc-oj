package main

import (
	_ "beegoDemo01/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//beego.SetStaticPath()
	beego.Run(":8888")
}

package rpc

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Panda int

/*
方法是可以导出的
方法要具有两个参数都是导出类型或者内建类型
方法的第二个参数必须是指针
方法只有一个返回值即error接口类型
*/

// Getinfo 函数关键字（对象）函数名(对端发送内容，返回对端内容)返回类型
func (this *Panda) Getinfo(argType int, replyType *int) error {
	fmt.Println("打印对端端口发送内容：", argType)
	*replyType = argType + 1
	return nil
}

func Server() {
	//	服务端必须注册对象 类实例化为对象
	pd := new(Panda)
	rpc.Register(pd)
	// 连接到网络
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("服务端启动失败err:", err)
		return
	}
	http.Serve(listener, nil)
}

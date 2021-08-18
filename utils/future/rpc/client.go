package rpc

import (
	"fmt"
	"net/rpc"
)

//rpc的客户端

func Client() {
	//建立网络连接
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("网络连接失败...")
		return
	}
	var pd int
	err = cli.Call("Panda.Getinfo", 1000, &pd)
	if err != nil {
		fmt.Println("调用失败...")
		return
	}
	fmt.Println("调用成功,pd:", pd)
}

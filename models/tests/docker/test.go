package main

import (
	"fmt"
	"time"
)

func main() {
	//打印当前时间 ,格式“2006-01-02 15:04:05 Mon
	//“2006-01-02 15:04:05是go的诞生时间，所以这么设计Format的Layout”
	for {
		var goos = time.Now().Format("2006-01-02 15:04:05")
		fmt.Printf("the current time:%s \n", goos)
		time.Sleep(1 * time.Second)
	}
}

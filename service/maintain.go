package service

//维护项目结构

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
	"time"
)

// StartCronTask 启动定时任务删除文件
func StartCronTask(validTime int) *cron.Cron {
	//定期 删除过期文件
	cron2 := cron.New() //创建一个cron实例
	//validTime秒执行一次
	err := cron2.AddFunc("*/"+strconv.Itoa(validTime)+" * * * * *", RemoveOutDatedFiles)
	if err != nil {
		logger.Error(err)
	}
	logger.Debug("定时删除任务启动成功...")
	cron2.Start()
	return cron2
}

// GetAllFile 递归获取指定目录下的所有文件名
func GetAllFile(pathname string) ([]string, error) {
	result := []string{}

	fis, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Printf("读取文件目录失败，pathname=%v, err=%v \n", pathname, err)
		return result, err
	}

	// 所有文件/文件夹
	for _, fi := range fis {
		fullname := pathname + "/" + fi.Name()
		// 是文件夹则递归进入获取;是文件，则压入数组
		if fi.IsDir() {
			temp, err := GetAllFile(fullname)
			if err != nil {
				fmt.Printf("读取文件目录失败,fullname=%v, err=%v", fullname, err)
				return result, err
			}
			result = append(result, temp...)
		} else {
			result = append(result, fullname)
		}
	}

	return result, nil
}

// SecondToTime 把秒级的时间戳转为time格式
func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// RemoveOutDatedFiles 删除codesfiles当中的过期文件
func RemoveOutDatedFiles() {
	// 递归获取目录下的所有文件
	var files []string
	files, _ = GetAllFile("../codesfiles")
	//fmt.Println("目录下的所有文件如下")
	for i := 0; i < len(files); i++ {
		//fmt.Println("文件名：", files[i])

		// 获取文件原来的访问时间，修改时间
		finfo, _ := os.Stat(files[i])

		// linux环境下代码如下
		linuxFileAttr := finfo.Sys().(*syscall.Stat_t)
		//fmt.Println("文件创建时间", SecondToTime(linuxFileAttr.Ctim.Sec))
		//fmt.Println("最后访问时间", SecondToTime(linuxFileAttr.Atim.Sec))
		//fmt.Println("最后修改时间", SecondToTime(linuxFileAttr.Mtim.Sec))
		//10分钟过期
		if time.Now().Sub(SecondToTime(linuxFileAttr.Mtim.Sec)) > time.Second*10 {
			err := os.Remove(files[i])
			if err != nil {
				logger.Error("文件删除失败：" + err.Error())
			}
			logger.Info("文件:" + files[i] + " 已过期，予以删除")
		}
		// windows下代码如下
		//winFileAttr := finfo.Sys().(*syscall.Win32FileAttributeData)
		//fmt.Println("文件创建时间：",SecondToTime(winFileAttr.CreationTime.Nanoseconds()/1e9))
		//fmt.Println("最后访问时间：",SecondToTime(winFileAttr.LastAccessTime.Nanoseconds()/1e9))
		//fmt.Println("最后修改时间：",SecondToTime(winFileAttr.LastWriteTime.Nanoseconds()/1e9))
	}
}

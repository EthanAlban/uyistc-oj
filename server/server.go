package server

import (
	"bufio"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/wonderivan/logger"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
	"unioj-Judger/conf"
	"unioj-Judger/public"
)

//这个文件用于编译运行代码

type Judger struct {
	TimeLimit   time.Duration //毫秒时间
	MemoryLimit int           //内存限制的单位是mb
	Code        string        //代码
	Language    string        //语言类型
	Suffix      string        //代码后缀
	FinalStat   int           //判题结束时的状态
	JudgeInfo   string        //判题得到的信息
}

func NewServer() *Judger {
	return &Judger{}
}

var JUDGER *Judger = &Judger{}

var wg sync.WaitGroup

func (*Judger) StartTask(result chan public.JudgeResult) {
	cfg := conf.CFG
	codesDir := cfg.Section("codesFile").Key("codesDir").String()
	filename := (uuid.NewV4()).String() + "." + JUDGER.Suffix
	JUDGER.CmdBash(codesDir, filename, result)
	ret := public.JudgeResult{Code: JUDGER.FinalStat, Info: JUDGER.JudgeInfo}
	result <- ret
}

func (*Judger) CmdBash(codesDir, filename string, result chan public.JudgeResult) {
	logger.Info(codesDir, "   +    ", filename)
	//fmt.Println(codesDir, "   +    ", filename)
	filepath := codesDir + filename
	_, err := JUDGER.CreateFile(filepath)
	if err != nil {
		return
	}
	err = JUDGER.WriteIntoFile(filepath, JUDGER.Code)
	if err != nil {
		return
	}
	err = JUDGER.CompileSourceFile(filepath)
	if err != nil {
		return
	}
	in := make(chan string, 1)
	out := make(chan string)
	timeout := make(chan bool, 1)
	//开启协程执行代码
	go JUDGER.ExecuteObj(filepath, in, out, timeout)
	//读取testcase的in的数组
	testcases_in := []string{"1 1", "2 1", "3 1", "4 1", "5 1", "6 1", "7 1", "8 1", "9 1"}
	testcases_out := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10"}
	//解答的结果中正确和错误的个数
	correct, wrong := 0, 0
	//in <- strconv.Itoa(0)
	for i := 0; i < len(testcases_in); i++ {
		in <- testcases_in[i]
		fmt.Println("1. 发送:", i+1)
		meminfo, err := GetProcessInfo(strings.Split(filename, ".")[0])
		if err == nil {
			//fmt.Println("rss:", meminfo.RSS/1024, "Mb  MemoryLimit:", JUDGER.MemoryLimit, "Mb")
			logger.Info("rss:", meminfo.RSS/1024, "Mb  MemoryLimit:", JUDGER.MemoryLimit, "Mb")
			if meminfo.RSS/1024 > JUDGER.MemoryLimit {
				JUDGER.JudgeInfo = ""
				JUDGER.FinalStat = 3
				break
			}
		}
		select {
		case num := <-out:
			fmt.Printf("3. ---------输出结果:%v  预期结果：%v\n", num, testcases_out[i])
			if num == testcases_out[i] {
				correct++
			} else {
				wrong++
			}
		case <-time.After(JUDGER.TimeLimit): //上面的ch如果一直没数据会阻塞，那么select也会检测其他case条件，检测到后3秒超时
			logger.Warn("用户程序计算超时" + strconv.Itoa(int(JUDGER.TimeLimit)/1000000) + "ms...")
			logger.Debug("超时准备干掉超时进程...")
			cmd := exec.Command("sh", "-c", "kill -9 `ps -aux | grep codesfiles | awk '{print $2}'`")
			_, err := cmd.CombinedOutput()
			if err != nil {
				logger.Error(err)
			}
			//fmt.Println("超时")
			JUDGER.JudgeInfo = ""
			JUDGER.FinalStat = 2
			i = len(testcases_in)
			//通道数据接收方不关闭通道
			timeout <- true
			close(in)
			//close(out)
			break
		}
	}
	//判题结果
	//全错了没有一个对的
	logger.Info("最终判题结果，correct:", correct, "  wrong:", wrong)
	//fmt.Println("最终判题结果，correct:", correct, "  wrong:", wrong)
	if correct == 0 && wrong == len(testcases_out) {
		JUDGER.JudgeInfo = ""
		JUDGER.FinalStat = -1
	} else if wrong == 0 && correct == len(testcases_out) {
		JUDGER.FinalStat = 0
	} else if wrong >= 0 && correct > 0 {
		JUDGER.FinalStat = 7
	}
	ret := public.JudgeResult{Code: JUDGER.FinalStat, Info: ""}
	result <- ret
}

func (*Judger) CreateFile(filepath string) (*os.File, error) {
	fp, err := os.Create("../" + filepath) // 如果文件已存在，会将文件清空。
	if err != nil {
		logger.Error("[1] 文件创建失败...,err:", err)
		//fmt.Println("[1] 文件创建失败...,err:", err)
		//创建文件失败的原因有：
		//1、路径不存在  2、权限不足  3、打开文件数量超过上限  4、磁盘空间不足等
		return fp, err
	}
	// defer延迟调用
	defer fp.Close() //关闭文件，释放资源。
	logger.Info("[1] 文件创建成功...")
	//fmt.Println("[1] 文件创建成功...")
	return fp, err
}

func (*Judger) WriteIntoFile(filepath, filecontent string) error {
	filePath := filepath
	file, err := os.OpenFile("../"+filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logger.Error("[2] 文件打开失败...", err)
		//fmt.Println("[2] 文件打开失败...", err)
		return err
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = write.WriteString(filecontent)
	if err != nil {
		return err
	}
	//Flush将缓存的文件真正写入到文件中
	err = write.Flush()
	if err != nil {
		return err
	}
	logger.Info("[2] 文件写入成功...")
	//fmt.Println("[2] 文件写入成功...")
	return err
}

func (*Judger) CompileSourceFile(filepath string) error {
	//cmd := exec.Command("ls", "-l", "/var/log/")
	//按照语言选择不同的编译运行命令
	var cmd *exec.Cmd
	logger.Debug(JUDGER.Language + " 代码，编译中...")
	logger.Error("-----------", filepath)
	if JUDGER.Language == "C" {
		cmd = exec.Command("gcc", "../"+filepath, "-o", "../"+strings.Split(filepath, ".")[0])
	} else if JUDGER.Language == "Golang" {
		cmd = exec.Command("go", "build", "-o", "../"+strings.Split(filepath, ".")[0], "../"+filepath)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		JUDGER.JudgeInfo = strings.Replace(string(out), filepath, "submissionFile ", -1)
		JUDGER.FinalStat = -2
		logger.Warn("[3] 编译文件失败:%s\n", string(out))
		//fmt.Printf("[3] 编译文件失败:%s\n", string(out))
		return err
	}
	logger.Info("[3] 编译文件成功... %s\n", string(out))
	//fmt.Printf("[3] 编译文件成功... %s\n", string(out))
	return err
}

func (*Judger) ExecuteObj(filepath string, inChan chan string, outChan chan string, timeout chan bool) {
	for {
		select {
		case num := <-inChan: //如果有数据，下面打印。但是有可能ch一直没数据
			cmd := exec.Command("../" + strings.Split(filepath, ".")[0])
			fmt.Println("2. 接收到：", num)
			cmd.Stdin = strings.NewReader(num)
			out, err := cmd.CombinedOutput()
			if err != nil {
				JUDGER.JudgeInfo = strings.Replace(string(out), filepath, "submissionFile ", -1) + err.Error()
				JUDGER.FinalStat = 4
				logger.Warn("[4] 执行文件失败...,combined out:%s err:%s\n", string(out), err)
				//fmt.Printf("[4] 执行文件失败...,combined out:%s err:%s\n", string(out), err)
			}
			outChan <- string(out)

			//fmt.Printf("[4] 执行文件成功,程序输出:%s\n", string(out))
		case delaty := <-timeout:
			fmt.Println("-+-+_++_+_+_+", delaty)
			close(outChan)
			return
		}
	}
}

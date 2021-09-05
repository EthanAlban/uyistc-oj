package server

//获取进程运行内存 判断是否超过限定内存
import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type MemoryStat struct {
	USER     string  //线程拥有者
	PID      string  //pid
	CpuUsage float64 //占用的 CPU 使用率
	MemUsage float64 //占用的记忆体使用率
	VSZ      int     //占用的虚拟记忆体大小  单位：kb（killobytes）
	RSS      int     //占用的记忆体大小   单位：kb（killobytes）
	TTY      int     //终端的次要装置号码 (minor device number of tty)
	STAT     string  //该行程的状态，linux的进程有5种状态： D 不可中断 uninterruptible sleep (usually IO)  R 运行 runnable (on run queue)  S 中断 sleeping  T 停止 traced or stopped   Z 僵死 a defunct (”zombie”) process
	START    string  //行程开始时间
	TIME     string  //执行的时间
	COMMAND  string  //所执行的指令
}

// GetProcessInfo 获取某个运行进程的信息
func GetProcessInfo(uuid string) (MemoryStat, error) {
	cmd := exec.Command("bash", "-c", "ps aux | grep "+uuid)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("获取进程信息失败...:%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
		return MemoryStat{}, err
	}
	res := DeleteExtraSpace(strings.Split(string(out), "\n")[0])
	resArr := strings.Split(res, " ")
	var processInfos MemoryStat
	processInfos.USER = resArr[0]
	processInfos.PID = resArr[1]
	processInfos.CpuUsage, _ = strconv.ParseFloat(resArr[2], 64)
	processInfos.MemUsage, _ = strconv.ParseFloat(resArr[3], 64)
	processInfos.VSZ, _ = strconv.Atoi(resArr[4])
	processInfos.RSS, _ = strconv.Atoi(resArr[5])
	processInfos.TTY, _ = strconv.Atoi(resArr[6])
	processInfos.STAT = resArr[7]
	processInfos.START = resArr[8]
	processInfos.TIME = resArr[9]
	processInfos.COMMAND = resArr[10]
	fmt.Printf("获取进程信息成功... \n%vMb\n", processInfos.RSS/1024)
	return processInfos, err
}

func Process() {
	fmt.Println(process.Pids())
	fmt.Println(process.MemoryMapsStat{}.String())
}

func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

package utils

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"unioj/models"
)

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"` //验证码Id
	ImageUrl  string `json:"imageUrl"`  //验证码图片url
}

func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// VerifyMobileFormat mobile verify
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

type Ret struct {
	Code int
	Data interface{}
	Msg  string
}

// CheckJudgerHealth 判断judgeserver的健康状态
func CheckJudgerHealth() {
	judgers := *(models.NewJudger().GetAllJudger())
	//judgers := strings.Split(conf.GetStringConfig("judgeserver_hosts"), ",")
	logger.Debug("[6] INFO judger健康检测启动...")
	for i := 0; i < len(judgers); i++ {
		Url, err := url.Parse("http://" + judgers[i].JudgerHost + "/healthy")
		if err != nil {
			logger.Error(err)
		}
		urlPath := Url.String()
		resp, err := http.Get(urlPath)
		if err != nil {
			logger.Error("	_" + strconv.Itoa(i) + "_   ERROR  judger:" + judgers[i].JudgerHost + "健康检测异常...\n  		err:" + err.Error())
			err := models.NewJudger().UpdateJudgerStatus(judgers[i].JudgerHost, 0, 0, 0)
			if err != nil {
				logger.Error("更新judger"+judgers[i].JudgerHost+"状态失败", err)
			}
			continue
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var ret Ret
		json.Unmarshal(body, &ret)
		if ret.Code == 200 {
			logger.Debug("	_" + strconv.Itoa(i) + "_   INFO  judger:" + judgers[i].JudgerHost + "健康检测正常...   ")
			//拿好的机器的参数
			mem, cpu := GetJudgerStatus(judgers[i].JudgerHost)
			logger.Debug(mem, cpu)
			err = models.NewJudger().UpdateJudgerStatus(judgers[i].JudgerHost, 1, mem, cpu)
			if err != nil {
				logger.Error("更新judger"+judgers[i].JudgerHost+"状态失败", err)
			}
		}
	}
}

type Usage struct {
	MemLoad float64 `json:"MemLoad"`
	LoadAvg float64 `json:"LoadAvg"`
}

func GetJudgerStatus(host string) (mem, cpu float64) {
	Url, err := url.Parse("http://" + host + "/usage_status")
	if err != nil {
		logger.Error(err)
		//logger.Error(err)
	}
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var ret Usage
	err = json.Unmarshal(body, &ret)
	if err != nil {
		logger.Error(err)
	}
	return ret.MemLoad, ret.LoadAvg
}

// StartJudgerHealtyCheck 定时检测健康状态
func StartJudgerHealtyCheck(validTime int) *cron.Cron {
	//定期 删除过期文件
	cron2 := cron.New() //创建一个cron实例
	//validTime秒执行一次
	err := cron2.AddFunc("*/"+strconv.Itoa(validTime)+" * * * * *", CheckJudgerHealth)
	if err != nil {
		logger.Error(err)
	}
	logger.Debug("定时检测判题器健康状态启动成功...")
	cron2.Start()
	return cron2
}

// PrintSysLogo 打印系统图形
func PrintSysLogo() {
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, "  ___       ___     ___         ___     ___     _____________     _______________\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |   |     |   |   |  \\  \\     |   |   |   |   |   _______   |   |_______________|\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |   |     |   |   |   \\  \\    |   |    ———    |  |       |  |         |  |\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |   |     |   |   |   |\\  \\   |   |    ___    |  |       |  |         |  |\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |   |     |   |   |   | \\  \\  |   |   |   |   |  |       |  |         |  |\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |   |     |   |   |   |  \\  \\ |   |   |   |   |  |       |  |         |  |\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |   |_____|   |   |   |   \\  \\|   |   |   |   |  |_______|  |     | | |  |\n", 0x1B)
	fmt.Printf("%c[1;5;34m%s%c[0m", 0x1B, " |_____________|   |___|    \\______|   |___|   |_____________|        |  |      V 1.0.0\n", 0x1B)
}

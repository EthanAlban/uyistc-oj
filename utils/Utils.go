package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"unioj/conf"
	logger "unioj/logs"
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
func CheckJudgerHealth() (err error) {
	judgers := strings.Split(conf.GetStringConfig("judgeserver_hosts"), ",")
	fmt.Println("[6] INFO judger健康检测启动...")
	for i := 0; i < len(judgers); i++ {
		Url, err := url.Parse("http://" + judgers[i] + "/healthy")
		if err != nil {
			return err
		}
		urlPath := Url.String()
		resp, err := http.Get(urlPath)
		if err != nil {
			//fmt.Println("[6] ERROR  judger:"+judgers[i]+"健康检测异常...   err:",err)
			logger.LogError("	_" + strconv.Itoa(i) + "_   ERROR  judger:" + judgers[i] + "健康检测异常...\n  		err:" + err.Error())
			continue
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		var ret Ret
		json.Unmarshal(body, &ret)
		if ret.Code == 200 {
			fmt.Println("	_" + strconv.Itoa(i) + "_   INFO  judger:" + judgers[i] + "健康检测正常...   ")
		}
	}
	return err
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

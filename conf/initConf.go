package conf

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Config 从beego服务器拉取配置文件
type Config struct {
	Data struct {
		JugerTaskTopic string `json:"juger_task_topic"`
		KafkaHost      string `json:"kafka_host"`
		SqlHost        string `json:"sql_host"`
		SqlUser        string `json:"sql_user"`
		SqlPassword    string `json:"sql_password"`
	} `json:"data"`
	Errcode int    `json:"errcode"`
	Message string `json:"message"`
}

var CFG *ini.File

func GetServerConfig() (*Config, error) {
	cfg, err := ini.Load("../conf/conf.ini")
	CFG = cfg
	if err != nil {
		logger.Error("Fail to read file: ", err)
	}
	serverHost := cfg.Section("server").Key("host").String()
	serverToken := cfg.Section("server").Key("token").String()
	//md5加密token
	data := []byte(serverToken)
	has := md5.Sum(data)
	md5token := fmt.Sprintf("%x", has)
	//设置参数
	params := url.Values{}
	Url, err := url.Parse("http://" + serverHost + "/api/utils/get_server_config")
	if err != nil {
		return nil, err
	}
	params.Set("token", md5token)
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		logger.Error("[1] ERROR  获取server kafka配置信息失败... ", err)
		logger.Error("[1] ERROR  请检查与服务器的连接...")
		//fmt.Println("[1] ERROR  获取server kafka配置信息失败...")
		//fmt.Println("[1] ERROR  请检查与服务器的连接...")
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var conf Config
	json.Unmarshal(body, &conf)
	if conf.Errcode != 200 {
		//如果出现这个错误请注意检查server/config中的token是否与服务器的一致
		logger.Error("[1] ERROR  获取server kafka配置信息失败...   err:", conf.Message)
		//fmt.Println("[1] ERROR  获取server kafka配置信息失败...   err:", conf.Message)
		return nil, errors.New(conf.Message)
	}
	cfg.Section("server").NewKey("juger_task_topic", conf.Data.JugerTaskTopic)
	cfg.Section("server").NewKey("kafka_host", conf.Data.KafkaHost)
	cfg.Section("server").NewKey("sql_host", conf.Data.SqlHost)
	cfg.Section("server").NewKey("sql_user", conf.Data.SqlUser)
	cfg.Section("server").NewKey("sql_password", conf.Data.SqlPassword)
	logger.Debug("[1] INFO  获取server kafka配置信息成功...")
	//fmt.Println("[1] INFO  获取server kafka配置信息成功...")
	logger.Info("		(_1_)  kafka_host:      ", cfg.Section("server").Key("kafka_host").String())
	logger.Info("		(_2_)  juger_task_topic:", cfg.Section("server").Key("juger_task_topic").String())
	logger.Info("		(_3_)  sql_host:        ", cfg.Section("server").Key("sql_host").String())
	logger.Info("		(_4_)  sql_user:        ", cfg.Section("server").Key("sql_user").String())
	logger.Info("		(_5_)  sql_password:    ", cfg.Section("server").Key("sql_password").String())
	return &conf, nil
}

package controllers

import (
	"encoding/json"
	"fmt"
	_ "github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
	"unioj/conf"
	"unioj/logs"
	"unioj/models"
	"unioj/models/redisOP"
)

type UtilsController struct {
	BaseController
}

type Weather struct {
	Result struct {
		HeWeather5 []struct {
			Aqi struct {
				City struct {
					Aqi  string `json:"aqi"`
					Co   string `json:"co"`
					No2  string `json:"no2"`
					O3   string `json:"o3"`
					Pm10 string `json:"pm10"`
					Pm25 string `json:"pm25"`
					Qlty string `json:"qlty"`
					So2  string `json:"so2"`
				} `json:"city"`
			} `json:"aqi"`
			Basic struct {
				City   string `json:"city"`
				Cnty   string `json:"cnty"`
				Id     string `json:"id"`
				Lat    string `json:"lat"`
				Lon    string `json:"lon"`
				Update struct {
					Loc string `json:"loc"`
					Utc string `json:"utc"`
				} `json:"update"`
			} `json:"basic"`
			DailyForecast []struct {
				Astro struct {
					Mr string `json:"mr"`
					Ms string `json:"ms"`
					Sr string `json:"sr"`
					Ss string `json:"ss"`
				} `json:"astro"`
				Cond struct {
					CodeD string `json:"code_d"`
					CodeN string `json:"code_n"`
					TxtD  string `json:"txt_d"`
					TxtN  string `json:"txt_n"`
				} `json:"cond"`
				Date string `json:"date"`
				Hum  string `json:"hum"`
				Pcpn string `json:"pcpn"`
				Pop  string `json:"pop"`
				Pres string `json:"pres"`
				Tmp  struct {
					Max string `json:"max"`
					Min string `json:"min"`
				} `json:"tmp"`
				Vis  string `json:"vis"`
				Wind struct {
					Deg string `json:"deg"`
					Dir string `json:"dir"`
					Sc  string `json:"sc"`
					Spd string `json:"spd"`
				} `json:"wind"`
			} `json:"daily_forecast"`
			Now struct {
				Cond struct {
					Code string `json:"code"`
					Txt  string `json:"txt"`
				} `json:"cond"`
				Fl   string `json:"fl"`
				Hum  string `json:"hum"`
				Pcpn string `json:"pcpn"`
				Pres string `json:"pres"`
				Tmp  string `json:"tmp"`
				Vis  string `json:"vis"`
				Wind struct {
					Deg string `json:"deg"`
					Dir string `json:"dir"`
					Sc  string `json:"sc"`
					Spd string `json:"spd"`
				} `json:"wind"`
			} `json:"now"`
			Status string `json:"status"`
		} `json:"HeWeather5"`
	} `json:"result"`
}

func (this *UtilsController) WetherToday() {
	weather_ := this.Ctx.Input.CruSession.Get("weather")
	if weather_ != nil {
		this.JsonResult(200, "缓存加载", weather_)
	}
	params := url.Values{}
	appkey := conf.GetStringConfig("appkey")
	city := "chengdu"
	Url, err := url.Parse("https://way.jd.com/he/freeweather")
	if err != nil {
		return
	}
	params.Set("city", city)
	params.Set("appkey", appkey)
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	//fmt.Println(urlPath) // https://httpbin.org/get?age=23&name=zhaofan
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var weather Weather
	json.Unmarshal(body, &weather)
	//记录session
	if this.Ctx.Input.CruSession == nil {
		this.StartSession()
	}
	this.Ctx.Input.CruSession.Set("weather", weather)
	this.JsonResult(200, "新查询", weather)
}

func (this *UtilsController) Heartbeat() {
	timeout := time.Duration(5 * time.Second)
	_, err := net.DialTimeout("tcp", "127.0.0.1:"+conf.GetStringConfig("httpport"), timeout)
	if err != nil {
		fmt.Println("Site unreachable, error: ", err)
		logs.LogError("服务器当前不可达...")
		this.JsonResult(204, "server not found", err)
	}
	this.JsonResult(200, "")
}

var cpt *captcha.Captcha

// Captcha 获取验证码
func (this *UtilsController) Captcha() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}

// RecordSysInfo 记录用户系统信息操作
func (this *UtilsController) RecordSysInfo() {
	userid := this.Ctx.Input.CruSession.Get("userid")
	if userid == nil {
		this.JsonResult(205, "未登录")
	}
	userid = userid.(int)
}

//
func (this *UtilsController) GetAllTags() {
	var tags []models.Tags
	tags_str, err := redisOP.RedisGetKey("Problemtags")
	if err == nil {
		json.Unmarshal([]byte(tags_str), &tags)
		this.JsonResult(200, "redis获取tags", tags)
	}
	tags_, err := models.NewTags().GetAllTags()
	tagsByte, _ := json.Marshal(tags_)
	err = redisOP.RedisSetKey("Problemtags", string(tagsByte))
	if err != nil {
		this.JsonResult(205, "缓存tags失败")
	}
	this.JsonResult(200, "获取tags成功", tags_)
}

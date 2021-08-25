// Copyright 2011 Dmitry Chestnykh. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// example of HTTP server that uses the captcha package.
package captcha

import (
	"encoding/json"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/wonderivan/logger"
	"net/http"
	"unioj/conf"
	//"text/template"
)

//var formTemplate = template.Must(template.New("example").Parse(formTemplateSrc))

type JsonResult struct {
	Errcode int    `json:"errcode"`
	Msg     string `json:"msg"`
}

func showFormHandler(w http.ResponseWriter, r *http.Request) {
	d := struct {
		CaptchaId string
	}{
		captcha.NewLen(4),
	}
	//if err := formTemplate.Execute(w, &d); err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	var id = make(map[string]string)
	id["CaptchaId"] = d.CaptchaId
	idStr, err := json.Marshal(id)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(idStr)
	//fmt.Println("新的验证码ID:"+d.CaptchaId)
}

func processFormHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var result = make(map[string]interface{})
	if !captcha.VerifyString(r.FormValue("captchaId"), r.FormValue("captchaSolution")) {
		result["errcode"] = 205
		result["msg"] = "Wrong captcha solution! No robots allowed!"
	} else {
		result["errcode"] = 200
		result["msg"] = "Great job, human! You solved the captcha."
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		logger.Error("验证验证码序列化失败：" + err.Error())
	}
	w.Write(bytes)
}

func StartCaptchaServer() {
	http.HandleFunc("/captcha/new", showFormHandler) //创建新的验证码
	http.HandleFunc("/captcha/verify", processFormHandler)
	http.Handle("/captcha/", captcha.Server(captcha.StdWidth*0.7, captcha.StdHeight))
	captchaServer := conf.GetStringConfig("captcha_host")
	logger.Debug("[4] INFO 验证码服务器启动... Server is at " + captchaServer)
	if err := http.ListenAndServe(captchaServer, nil); err != nil {
		logger.Error("验证码服务器启动失败：" + err.Error())
	}
}

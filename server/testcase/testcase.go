package testcase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Testcase struct {
	Result []struct {
		Output     string `json:"output"`
		Input      string `json:"input"`
		TestcaseId int64  `json:"testcase_id"`
	} `json:"result"`
}

func GetProblemTestcases(pid int) ([]string, []string, []float64) {
	params := url.Values{}
	//cfg := conf.CFG
	//host := cfg.Section("filesever").Key("host").String()
	host := "http://127.0.0.1:8081"
	Url, err := url.Parse(host + "/question/getCaseMsgByPid")
	if err != nil {
		fmt.Println(err)
		return nil, nil, nil
	}
	params.Set("pid", strconv.Itoa(pid))
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var testcase Testcase
	json.Unmarshal(body, &testcase)
	input := make([]string, 0)
	output := make([]string, 0)
	scores := make([]float64, len(testcase.Result))
	for _, case_ := range testcase.Result {
		input = append(input, case_.Input)
		output = append(output, case_.Output)
	}
	for i := 0; i < len(testcase.Result); i++ {
		scores[i] = float64(100 / len(testcase.Result))
	}
	return input, output, scores
}

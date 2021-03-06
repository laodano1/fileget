package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davyxu/golog"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

type MarqueeRulesRecords struct {
	Records []*MarqueeRulesResp `json:"records"`
}

type PostResultNew struct {
	Code    int         `json:"code"`    //200:成功 详见【状态码】
	Message string      `json:"message"` //接口描述信息
	Data    interface{} `json:"data"`    //请求json数据
}

type MarqueeRulesResp struct {
	GameId           string `json:"gameId"`      //游戏id
	AgentId          int32  `json:"agentId"`     //业主编号
	RoomId           string `json:"roomId"`      //房间编号
	AmountLimit      int64  `json:"amountLimit"` //金额限制
	Content          string `json:"content"`
	RuleId           int64  `json:"ruleId"`
	SpecialCondition string `json:"specialCondition"` //
}

// 获取跑马灯规则
type GetMarqueeRulesCommand struct {
	GameId string `json:"gameId"` //游戏id
}

const (
	userAgent = "go client"
	//contentType = "application/json; charset=utf-8"
	timeOut = time.Second * 2
)

var contentType = "application/json"

func PPost(url string, body []byte, agentId int32, timeoutReq ...time.Duration) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(resp)
	}()
	req.SetRequestURI(url)
	req.Header.SetContentType(contentType)
	req.Header.SetUserAgent(userAgent)
	req.Header.SetMethod(http.MethodPost)

	if agentId == 0 {
		//req.Header.Add("Agent", "0")
	} else {
		req.Header.Add("Agent", fmt.Sprintf("%v", agentId))
	}
	fmt.Printf("============================ http request: %v\n", req)
	req.SetBody(body)
	if len(timeoutReq) == 0 {
		timeoutReq = append(timeoutReq, timeOut)
	}

	client := &fasthttp.Client{
		MaxConnsPerHost: 1024,
		ReadTimeout: timeOut,
		WriteTimeout: timeOut,
	}

	if err := client.DoTimeout(req, resp, timeoutReq[0]); err != nil {
		return nil, errors.New(fmt.Sprintf("调用服务%s超时：%s", url, err.Error()))
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("调用服务%s失败:%s", url, http.StatusText(resp.StatusCode())))
	}
	buf := new(bytes.Buffer)
	buf.Write(resp.Body())
	return buf.Bytes(), nil
}

func SearchMarqueeRules(gameId string, val interface{}) error {

	command := &GetMarqueeRulesCommand{
		GameId: gameId,
	}
	fmt.Printf("=================== SearchMarqueeRules: %v | %v\n", command, time.Now().Format(time.RFC3339Nano))
	err := postJsonNew("http://10.0.0.20:8080" + "/game/searchMarqueeRules", command, val, 0)
	if err != nil {
		return err
	}

	return nil
}


// new for platform
func postJsonNew(url string, data interface{}, val interface{}, agentId int32) (err error) {

	jsonStr, _ := json.Marshal(data)
	//req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req.Header.Add("content-type", "application/json")
	//if agentId == 0 {
	//	req.Header.Add("Agent", "0")
	//} else {
	//	req.Header.Add("Agent", fmt.Sprintf("%v", agentId))
	//}
	//if err != nil {
	//	return err
	//}
	//defer req.Body.Close()

	//
	//client := &http.Client{Timeout: 5 * time.Second}
	//fmt.Printf("============================ http request: %v\n", req)
	//resp, error := client.Do(req)
	//if error != nil {
	//	return error
	//}
	//defer resp.Body.Close()


	//fmt.Printf("============================ http request: %v\n", req)
	resp, err := PPost(url, jsonStr, agentId, timeOut)
	if err != nil {
		return err
	}

	result := &PostResultNew{}

	//var b []byte
	//b, err = ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return err
	//}

	//err = json.Unmarshal(b, result)
	err = json.Unmarshal(resp, result)
	if err != nil {
		return errors.New(fmt.Sprintf("%v | %v", err, string(resp)))
		//return errors.New(fmt.Sprintf("%v | %v", err, string(b)))
	}

	if 200 != result.Code {
		return errors.New(fmt.Sprintf("code(%v): %v", result.Code, result))
	}

	rd, err := json.Marshal(result.Data)
	if err != nil {
		return
	}

	if val != nil {
		err = json.Unmarshal(rd, val)
		if err != nil {
			return
		}
	}

	return
}

func main() {
	//bytes.Trim()
	lg := golog.New("test")
	mRuls := &MarqueeRulesRecords{}
	err := SearchMarqueeRules("600101", mRuls)
	if err != nil {
		lg.Errorf("SearchMarqueeRules failed: '%v'", err)
		return
	}

	lg.Debugf("marquee rules: %v", mRuls.Records)
}

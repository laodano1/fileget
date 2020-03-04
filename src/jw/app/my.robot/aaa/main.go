package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/davyxu/golog"
	"io/ioutil"
	"net/http"
	"time"
)

type PostResultNew struct {
	Code    int    `json:"code"`    //200:成功 详见【状态码】
	Message string `json:"message"` //接口描述信息
	Data    interface{} `json:"data"`    //请求json数据
}

type NotifyUserScopeData struct {
	GameId string 	`json:"gameId"`
	UserId int64    `json:"userId"`
	AgentId int32   `json:"agentId"`
	Balance int64   `json:"balance"`
	ChangeAmount int64   `json:"changeAmount"`
}


func postJsonNewNoRV(url string, data interface{}, agentId int32) (err error) {

	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", "application/json")
	if agentId == 0 {
		req.Header.Add("Agent", "")
	} else {
		req.Header.Add("Agent", fmt.Sprintf("%v", agentId))
	}
	if err != nil {
		return err
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		return error
	}
	defer resp.Body.Close()

	result := &PostResultNew{}

	var b []byte
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, result)
	if err != nil {
		return errors.New(fmt.Sprintf("%v | %v", err, string(b)))
	}

	if 200 != result.Code {
		return errors.New(result.Message)
	}


	return
}

type freezeUserNtf struct {
	AgentId int32 `json:"agentId"`
	UserId  int64 `json:"userId"`
	GameId  string `json:"gameId"`
}

type balanceChgNtf struct {
	Balance int64 `json:"balance"`
	GameId  string `json:"gameId"`
	UserId  int64 `json:"userId"`
	ChangeAmount int64 `json:"changeAmount"`
}

func main() {
	logger := golog.New("ttsst")

	encodedstr := base64.StdEncoding.EncodeToString([]byte("600101"))
	gameOfflinePostCmd := &PostResultNew{
		Code: 40809,
		Message: "游戏维护中，请稍后再试！",
		Data: encodedstr,
	}

	err := postJsonNewNoRV("http://10.0.0.37:8802" + "/game/notify", gameOfflinePostCmd, 0)
	if err != nil {
		logger.Errorf("error: %v", err)
		return
	}


	logger.Debugf("test")

	aun := &freezeUserNtf{
		AgentId: 80,
		UserId:  111,
		GameId: "600101",
	}

	data, _ := json.Marshal(aun)
	//encodedstr := base64.StdEncoding.EncodeToString([]byte("600101"))
	encodedstr40403 := base64.StdEncoding.EncodeToString(data)
	userFreezPostCmd := &PostResultNew{
		Code: 40403,
		Message: "您当前账号已被冻结无法登录，如有问题请联系客服！",
		Data: encodedstr40403,
	}
	err = postJsonNewNoRV("http://10.0.0.37:8802" + "/game/notify", userFreezPostCmd, 0)
	if err != nil {
		logger.Errorf("error: %v", err)
		return
	}

	bcn := &balanceChgNtf{
		Balance: 80,
		GameId: "600101",
		UserId:  111,
		ChangeAmount: 200,
	}

	data, _ = json.Marshal(bcn)
	//encodedstr := base64.StdEncoding.EncodeToString([]byte("600101"))
	encodedstr2005 := base64.StdEncoding.EncodeToString(data)
	bcnPostCmd := &PostResultNew{
		Code: 2005,
		Message: "您当前账号已被冻结无法登录，如有问题请联系客服！",
		Data: encodedstr2005,
	}
	err = postJsonNewNoRV("http://10.0.0.37:8802" + "/game/notify", bcnPostCmd, 0)
	if err != nil {
		logger.Errorf("error: %v", err)
		return
	}


	logger.Debugf("time: %d", time.Now().UnixNano())

}

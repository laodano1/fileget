package main

import (
	"fileget/util"
	"github.com/valyala/fasthttp"
)

var payLoad = `{"bk_biz_id":2,"ip":{"data":[],"flag":"bk_host_innerip|bk_host_outerip","exact":0},"page":{"start":0,"limit":20,"sort":"bk_host_id"},"condition":[{"bk_obj_id":"biz","fields":[],"condition":[{"field":"bk_biz_id","operator":"$eq","value":2}]},{"bk_obj_id":"set","fields":["bk_set_name","bk_set_id"],"condition":[]},{"bk_obj_id":"module","fields":["bk_module_name","bk_module_id"],"condition":[]},{"bk_obj_id":"host","fields":["bk_host_id","bk_host_innerip","bk_cloud_id","bk_host_outerip","operator"],"condition":[]},{"bk_obj_id":"object","fields":[],"condition":[]}]}`

// 蓝鲸文档中心 > 配置平台 > 查询实例关联拓扑
// https://bk.tencent.com/docs/document/5.1/9/382
func main() {

	urlStr := "https://www.baidu.com"
	urlStr = "https://cmdb.bk.tencent.com/api/v3/hosts/search"

	req := new(fasthttp.Request)
	rsp := new(fasthttp.Response)
	req.Header.SetMethod("POST")
	//req.Header.SetMethod("GET")
	req.SetRequestURI(urlStr)
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36")

	req.Header.Set("", "")

	//req.SetBody([]byte(payLoad))
	req.SetBodyString(payLoad)

	client := new(fasthttp.Client)
	if err := client.Do(req, rsp); err != nil {
		util.Lg.Errorf("error: %v", err)
		return
	}

	util.Lg.Debugf("%v", string(rsp.Body()))

}

package util

import (
	"github.com/astaxie/beego/httplib"
	"errors"
	"github.com/astaxie/beego"
	"time"
)

//httpclient 请求get util
func Get(url string, param map[string]interface{}) string {

	req := httplib.Get(url).SetTimeout(100 * time.Second, 100 * time.Second)
	req.Header("Accept-Encoding", "gzip,deflate,sdch")
	req.Header("Host", "beego.me")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")

	if len(param) != 0 {
		for key, _ := range param {
			req.Param(key, param[key])
		}
	}

	str, err := req.String()

	if err != nil {
		errors.New("获取请求数据异常!" + err)
		beego.Debug("http调用异常", err)
	}

	return str
}

//post 方法获取数据
func Post(url string, param map[string]interface{}) string {

	req := httplib.Post(url).SetTimeout(100 * time.Second, 100 * time.Second)
	req.Header("Accept-Encoding", "gzip,deflate,sdch")
	req.Header("Host", "beego.me")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")

	if len(param) != 0 {
		for key, _ := range param {
			req.Param(key, param[key])
		}
	}

	str, err := req.String()

	if (err != nil) {
		beego.Debug("http调用异常", err)
	}

	return str
}
//请求提交 json 数据
func RequestJson(url string, data string) string {

	req := httplib.Post(url).SetTimeout(100 * time.Second, 100 * time.Second)
	req.Header("Accept-Encoding", "gzip,deflate,sdch")
	req.Header("Host", "beego.me")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")

	req.JSONBody(data)

	str, err := req.String()
	if err != nil {
		beego.Debug("请求获取异常 json请求:", err)
	}
	return str
}





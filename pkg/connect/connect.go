package connect

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

// client 全局的HTTP客户端
var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: 2 * time.Second,
}

// Get 判断url是否能请求通
func Get(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		logx.Errorw("connect client.Get failed", logx.LogField{
			Key:   "merr",
			Value: err.Error(),
		})
	}
	err = resp.Body.Close()
	if err != nil {
		logx.Errorw("close body failed", logx.LogField{Key: "merr", Value: err.Error()})
		return false
	}
	return resp.StatusCode == http.StatusOK
}

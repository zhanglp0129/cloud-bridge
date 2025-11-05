package client

import (
	"bytes"
	"io"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/zhanglp0129/cloud-bridge/client-bridge/config"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRsp struct {
	Token string `json:"token"`
}

// Login 登录，返回token
func Login(username, password string) (string, error) {
	// 构造请求内容
	url := config.C.ServerURL + "/login"
	req := LoginReq{
		Username: config.C.Username,
		Password: config.C.Password,
	}
	body, err := sonic.Marshal(req)
	if err != nil {
		return "", err
	}
	// 发出请求
	rsp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()
	// 获取token
	rspBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	loginRsp := LoginRsp{}
	if err := sonic.Unmarshal(rspBytes, &loginRsp); err != nil {
		return "", err
	}
	return loginRsp.Token, nil
}

package client

import (
	"fmt"
	"io"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/zhanglp0129/cloud-bridge/client-bridge/config"
)

type ConfigRsp struct {
	RemoteName string `json:"remote_name"`
	RemoteType string `json:"remote_type"`
	Config     struct {
		RemotePath string `json:"remote_path"`
	} `json:"config"`
}

// RcloneConfig 获取rclone配置
func RcloneConfig(token string) (*ConfigRsp, error) {
	url := config.C.ServerURL + "/config/rclone"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+token)
	// 发出请求
	cli := http.Client{}
	rsp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	if rsp.StatusCode == 401 {
		return nil, fmt.Errorf("token error")
	}
	// 获取rclone配置
	rspBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	configRsp := ConfigRsp{}
	if err := sonic.Unmarshal(rspBytes, &configRsp); err != nil {
		return nil, err
	}
	return &configRsp, nil
}

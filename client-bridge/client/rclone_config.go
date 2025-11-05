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

const rcloneTomlTemplate string = `[%s]
type = %s
remote_path = %s
`

// RcloneConfig 获取rclone配置，返回rclone toml配置
func RcloneConfig(token string) (string, error) {
	url := config.C.ServerURL + "/config/rclone"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", token)
	// 发出请求
	cli := http.Client{}
	rsp, err := cli.Do(req)
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()
	// 获取rclone配置
	rspBytes, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	configRsp := ConfigRsp{}
	if err := sonic.Unmarshal(rspBytes, &configRsp); err != nil {
		return "", err
	}
	return fmt.Sprintf(rcloneTomlTemplate,
		configRsp.RemoteName,
		configRsp.RemoteType,
		configRsp.Config.RemotePath), nil
}

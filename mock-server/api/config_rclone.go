package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanglp0129/cloud-bridge/mock-server/util"
)

type ConfigRsp struct {
	RemoteName string `json:"remote_name"`
	RemoteType string `json:"remote_type"`
	Config     struct {
		RemotePath string `json:"remote_path"`
	} `json:"config"`
}

func ConfigRclone(c *gin.Context) {
	// 获取token
	auth := c.GetHeader("Authorization")
	if auth != "Bearer mock-jwt-token-12345" {
		// 认证失败
		c.String(401, "")
		return
	}
	config := ConfigRsp{}
	config.RemoteName = "MyCloudDrive"
	config.RemoteType = "local"
	config.Config.RemotePath = util.GetRemotePath()
	c.JSON(200, config)
}

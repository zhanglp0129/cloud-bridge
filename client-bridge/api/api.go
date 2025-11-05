package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanglp0129/cloud-bridge/client-bridge/client"
)

var rcloneConfigGlobal client.ConfigRsp

func NewApi(rcloneConfig *client.ConfigRsp) *gin.Engine {
	if rcloneConfig != nil {
		rcloneConfigGlobal = *rcloneConfig
	}
	r := gin.Default()
	r.GET("/health", Health)
	r.GET("/status", Status)
	r.POST("/trigger-sync", TriggerSync)
	return r
}

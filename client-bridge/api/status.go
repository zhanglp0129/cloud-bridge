package api

import (
	"os"

	"github.com/bytedance/sonic"
	"github.com/gin-gonic/gin"
	"github.com/zhanglp0129/cloud-bridge/client-bridge/config"
)

type StatusRsp struct {
	IsMounted  bool
	MountPoint string
}

func Status(c *gin.Context) {
	// 判断挂载点是否存在
	_, err := os.Stat(config.C.MountPath)
	isMounted := err == nil
	rspBytes, err := sonic.Marshal(StatusRsp{
		IsMounted:  isMounted,
		MountPoint: config.C.MountPath,
	})
	if err != nil {
		c.String(400, "marshal json error: %v", err)
	}
	c.JSON(200, string(rspBytes))
}

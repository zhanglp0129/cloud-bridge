package main

import (
	"log"

	"github.com/zhanglp0129/cloud-bridge/mock-server/api"
	"github.com/zhanglp0129/cloud-bridge/mock-server/util"
)

func main() {
	// 创建远程路径
	if err := util.CreateRemotePath(); err != nil {
		log.Fatalf("create remote path error: %v", err)
	}
	// 启动服务器
	s := api.NewApi()
	if err := s.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("run server error: %v", err)
	}
}

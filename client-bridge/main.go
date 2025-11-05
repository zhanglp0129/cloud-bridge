package main

import (
	"log"

	"github.com/zhanglp0129/cloud-bridge/client-bridge/config"
)

func main() {
	// 读取配置
	if err := config.Init(); err != nil {
		log.Fatalf("config init error: %v", err)
	}
}
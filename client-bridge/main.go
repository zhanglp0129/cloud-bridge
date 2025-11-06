package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/zhanglp0129/cloud-bridge/client-bridge/api"
	"github.com/zhanglp0129/cloud-bridge/client-bridge/client"
	"github.com/zhanglp0129/cloud-bridge/client-bridge/config"
	"github.com/zhanglp0129/cloud-bridge/client-bridge/util"
)

func main() {
	// 读取配置
	if err := config.Init(); err != nil {
		log.Fatalf("config init error: %v", err)
	}
	// 登录获取token
	token, err := client.Login(config.C.Username, config.C.Password)
	if err != nil {
		log.Fatalf("login error: %v", err)
	}
	// 获取rclone配置
	rcloneConfig, err := client.RcloneConfig(token)
	if err != nil {
		log.Fatalf("rclone config error: %v", err)
	}
	rcloneConfigToml := util.BuildRcloneConfigToml(rcloneConfig)
	// 写入配置到文件中
	if err := util.WriteRcloneConfig(rcloneConfigToml); err != nil {
		log.Fatalf("write rclone config error: %v", err)
	}
	// 启动挂载
	cmd, err := rcloneMount(rcloneConfig)
	if err != nil {
		log.Fatalf("rclone mount error: %v", err)
	}
	// 处理信号
	signalHandle(cmd)
	// 启动服务器
	s := api.NewApi(rcloneConfig)
	if err := s.Run("0.0.0.0:9527"); err != nil {
		unmount(cmd)
		log.Fatalf("server run error: %v", err)
	}
}

// rclone在后台挂载
func rcloneMount(rcloneConfig *client.ConfigRsp) (*exec.Cmd, error) {
	// 获取挂载路径
	mountPath := config.C.MountPath
	// 获取配置文件路径
	configFilename, err := util.GetRcloneConfigFilename()
	if err != nil {
		return nil, err
	}
	log.Printf("config file name: %s", configFilename)
	// 挂载命令
	args := []string{
		"mount",
		fmt.Sprintf("%s:", rcloneConfig.RemoteName),
		mountPath,
		"--config", configFilename,
		"--vfs-cache-mode", "full",
	}
	cmd := exec.Command("rclone", args...)
	log.Printf("command: %v", cmd.Args)
	// 设置工作目录
	cmd.Dir = rcloneConfig.Config.RemotePath
	// 捕获stdout和stderr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// 后台启动rclone进程
	go func() {
		if err := cmd.Run(); err != nil {
			log.Fatalf("rclone mount cmd run error: %v", err)
		}
	}()
	return cmd, nil
}

// 处理信号
func signalHandle(cmd *exec.Cmd) {
	signCh := make(chan os.Signal, 1)
	signal.Notify(signCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signCh
		log.Printf("received signal: %v", sig)
		unmount(cmd)
	}()
}

// 卸载
func unmount(cmd *exec.Cmd) {
	if cmd != nil && cmd.Process != nil {
		// 优雅地卸载
		unmountCmd := exec.Command("rclone", "unmount", config.C.MountPath)
		if err := unmountCmd.Run(); err != nil {
			log.Printf("run unmount error: %v", err)
		}
		// 如果还存在，则强制关闭
		if cmd.ProcessState == nil || !cmd.ProcessState.Exited() {
			cmd.Process.Kill()
		}
		os.Exit(0)
	}
}

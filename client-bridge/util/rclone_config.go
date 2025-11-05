package util

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/zhanglp0129/cloud-bridge/client-bridge/client"
)

// GetRcloneConfigFilename 获取rclone配置文件位置
func GetRcloneConfigFilename() (string, error) {
	// 获取家目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filename := path.Join(homeDir, ".config/client-bridge/rclone.conf")
	// 修正目录分隔符
	filename = strings.ReplaceAll(filename, "\\", "/")
	return filename, err
}

// WriteRcloneConfig 将rclone配置写入到文件中
func WriteRcloneConfig(rcloneConfig string) error {
	filename, err := GetRcloneConfigFilename()
	if err != nil {
		return err
	}
	log.Printf("rclone config filename: %v", filename)
	// 创建文件所在目录
	dir := path.Dir(filename)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	log.Printf("rclone config dirname: %v", dir)
	if err := os.WriteFile(filename, []byte(rcloneConfig), 0744); err != nil {
		return err
	}
	return nil
}

const rcloneTomlTemplate string = `[%s]
type = %s
remote_path = %s
`

// BuildRcloneConfigToml 构造rclone toml配置
func BuildRcloneConfigToml(rcloneConfig *client.ConfigRsp) string {
	return fmt.Sprintf(rcloneTomlTemplate,
		rcloneConfig.RemoteName,
		rcloneConfig.RemoteType,
		rcloneConfig.Config.RemotePath)
}

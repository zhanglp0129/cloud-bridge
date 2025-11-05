package config

import "github.com/spf13/viper"

var C Config

type Config struct {
	ServerURL string `json:"server_url"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	MountPath string `json:"mount_path"`
}

// Init 初始化配置
func Init() error {
	// 从配置文件读取
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	// 从环境变量读取
	viper.SetEnvPrefix("CLIENT_BRIDGE")
	viper.AutomaticEnv()
	
	if err := viper.Unmarshal(&C); err != nil {
		return err
	}
	return nil
}
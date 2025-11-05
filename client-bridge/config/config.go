package config

import (
	"log"

	"github.com/spf13/viper"
)

var C Config

type Config struct {
	ServerURL string `mapstructure:"server_url"`
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	MountPath string `mapstructure:"mount_path"`
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
	log.Printf("config: %+v", C)
	return nil
}

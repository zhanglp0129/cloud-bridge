package util

import "os"

func CreateRemotePath() error {
	// 检查目录是否存在
	if _, err := os.Stat(GetRemotePath()); err == nil {
		return nil
	}
	return os.Mkdir(GetRemotePath(), 0755)
}
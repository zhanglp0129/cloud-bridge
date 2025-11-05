package util

import "runtime"

func GetRemotePath() string {
	if runtime.GOOS == "windows" {
		return "C:\\Users\\zhang\\AppData\\Local\\Temp\\mock-remote-storage"
	}
	return "/tmp/mock-remote-storage"
}

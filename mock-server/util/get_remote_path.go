package util

import (
	"os"
	"path"
	"strings"
)

func GetRemotePath() string {
	path := path.Join(os.TempDir(), "mock-remote-storage")
	return strings.ReplaceAll(path, "\\", "/")
}

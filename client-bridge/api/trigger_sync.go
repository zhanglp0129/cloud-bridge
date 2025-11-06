package api

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/zhanglp0129/cloud-bridge/client-bridge/config"
	"github.com/zhanglp0129/cloud-bridge/client-bridge/util"
)

func TriggerSync(c *gin.Context) {
	configFile, err := util.GetRcloneConfigFilename()
	if err != nil {
		c.String(400, "get rclone config filename error: %v", err)
	}
	args := []string{
		"sync",
		fmt.Sprintf("%s:", rcloneConfigGlobal.RemoteName),
		config.C.MountPath,
		"--config", configFile,
	}
	cmd := exec.Command("rclone", args...)
	cmd.Dir = rcloneConfigGlobal.Config.RemotePath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		c.String(400, "rclone sync error: %v", err)
	}
	c.JSON(200, map[string]any{
		"status": "sync_triggered",
	})
}

package api

import "github.com/gin-gonic/gin"

func Login(c *gin.Context) {
	c.JSON(200, map[string]any{
		"token": "mock-jwt-token-12345",
	})
}
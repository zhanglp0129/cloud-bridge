package api

import "github.com/gin-gonic/gin"

func NewApi() *gin.Engine {
	r := gin.Default()
	r.POST("/login", Login)
	return r
}
package api

import "github.com/gin-gonic/gin"

// Setup returns configure gin router
func Setup() *gin.Engine {
	r := gin.New()
	return r
}

package logger

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// LogHandler simple JSON request logger
func LogHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Now().Sub(start)
		log.Debug().Str("ip", c.ClientIP()).Str("latency", latency.String()).Str("method", c.Request.Method).Str("uri", c.Request.RequestURI).Str("code", strconv.Itoa(c.Writer.Status())).Send()
	}
}

package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health handler returns database and telegram api status
func Health(c *gin.Context) {
	// FIXME: check mongo and telegram api status
	c.JSON(http.StatusOK, gin.H{"database": "ok", "telegram": "ok"})
}

//GetSettings handler returns lotus settings
func GetSettings(c *gin.Context) {
	c.Status(http.StatusNotImplemented)

}

// UpdateSettings handler updates lotus settings
func UpdateSettings(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

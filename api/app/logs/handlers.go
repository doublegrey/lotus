package logs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get handler returns app's logs
func Get(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

// Write handler creates new log record
func Write(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

// Delete handler clears app's logs
func Delete(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

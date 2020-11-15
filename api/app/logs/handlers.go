package logs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get handler returns app's logs
func Get(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

// Create handler creates new log record
func Create(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

// DeleteAll handler clears app's logs
func DeleteAll(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

// Delete handler deletes log by id
func Delete(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

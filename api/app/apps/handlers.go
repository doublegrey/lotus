package apps

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAll handler returns all apps
func GetAll(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

// Get handler returns app by id
func Get(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

// Create handler creates app
func Create(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

// Update handlers updates app settings
func Update(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

// Delete handler deletes app and its logs
func Delete(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

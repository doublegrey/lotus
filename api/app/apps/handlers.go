package apps

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/doublegrey/lotus/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetAll handler returns all apps
func GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "created", Value: -1}})
	cursor, err := utils.DB.Collection("apps").Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	var apps []App
	if err = cursor.All(ctx, &apps); err != nil {
		log.Fatal(err)
	}
	bytes, _ := json.Marshal(apps)
	c.JSON(http.StatusOK, gin.H{"data": string(bytes)})
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

package telegram

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/doublegrey/lotus/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateSettings handler updates telegram settings
func UpdateSettings(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var settings Settings
	err := c.ShouldBindJSON(&settings)
	if err != nil {
		log.Println(err)
	}
	settings.Name = "telegram"

	_, err = utils.DB.Collection("integrations").UpdateOne(ctx, bson.M{"name": "telegram"}, bson.M{"$set": settings}, options.Update().SetUpsert(true))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "app id is incorrect"})
		return
	}
	c.Status(http.StatusOK)
}

// GetSettings handler returns telegram settings
func GetSettings(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result := utils.DB.Collection("integrations").FindOne(ctx, bson.M{"name": "telegram"})
	var settings Settings
	err := result.Decode(&settings)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": settings})
}

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
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		log.Println(err)
	}
	var apps []App
	if err = cursor.All(ctx, &apps); err != nil {
		log.Println(err)
	}
	bytes, _ := json.Marshal(apps)
	c.JSON(http.StatusOK, gin.H{"data": string(bytes)})
}

// Get handler returns app by id
func Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var app App
	objectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Println()
	}
	result := utils.DB.Collection("apps").FindOne(ctx, bson.M{"_id": objectID})
	err = result.Decode(&app)
	if err != nil {
		log.Println(err)
	}
	bytes, _ := json.Marshal(app)
	c.JSON(http.StatusOK, gin.H{"data": string(bytes)})
}

// Create handler creates app
func Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var app App
	err := c.ShouldBindJSON(&app)
	if err != nil {
		log.Println(err)
	}
	app.ID = primitive.NewObjectIDFromTimestamp(time.Now().UTC())
	app.Created = time.Now().UTC()
	_, err = utils.DB.Collection("apps").InsertOne(ctx, app)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusCreated, gin.H{"id": app.ID})
}

// Update handlers updates app
func Update(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

// Delete handler deletes app and its logs
func Delete(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}

package apps

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/doublegrey/lotus/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Cache map contains all apps
var Cache sync.Map

// InitCache function initializes apps.Cache map
func InitCache() {
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
	for _, app := range apps {
		Cache.Store(app.ID, app)
	}
}

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
	c.JSON(http.StatusOK, gin.H{"data": apps})
}

// Get handler returns app by id
func Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var app App
	objectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Println(err)
	}
	result := utils.DB.Collection("apps").FindOne(ctx, bson.M{"_id": objectID})
	err = result.Decode(&app)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": app})
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
	Cache.Store(app.ID, app)
	c.JSON(http.StatusCreated, gin.H{"id": app.ID})
}

// Update handlers updates app
func Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var app App
	err := c.ShouldBindJSON(&app)
	if err != nil {
		log.Println(err)
	}
	if app.ID == primitive.NilObjectID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "app id is empty"})
		return
	}
	_, err = utils.DB.Collection("apps").UpdateOne(ctx, bson.M{"_id": app.ID}, bson.M{"$set": app})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "app id is incorrect"})
		return
	}
	Cache.Store(app.ID, app)
	c.JSON(http.StatusOK, gin.H{"id": app.ID})
}

// Delete handler deletes app and its logs
func Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Println(err)
	}
	_, err = utils.DB.Collection("apps").DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Println(err)
	}
	Cache.Delete(objectID)
	c.Status(http.StatusOK)
}

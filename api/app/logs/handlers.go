package logs

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/doublegrey/lotus/api/app/apps"
	"github.com/doublegrey/lotus/utils"
	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// VerifyToken middleware
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		appID, err := primitive.ObjectIDFromHex(c.Param("app"))
		if err != nil {
			log.Println(err)
		}
		a, exists := apps.Cache.Load(appID)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("failed to find app with id: %s", appID)})
			return
		}
		app := a.(apps.App)

		if len(app.Token) > 0 && token != app.Token {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "access token is empty or invalid"})
			return
		}
		c.Set("app", app)
		c.Next()
	}
}

// GetAll handler returns app's logs
func GetAll(c *gin.Context) {
	a, exists := c.Get("app")
	if !exists {
		c.Status(http.StatusInternalServerError)
		return
	}
	app := a.(apps.App)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "created", Value: -1}})
	cursor, err := utils.DB.Collection(app.Name).Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Println(err)
	}
	var records []Record
	if err = cursor.All(ctx, &records); err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": records})
}

// Get handler returns log record by id
func Get(c *gin.Context) {
	a, exists := c.Get("app")
	if !exists {
		c.Status(http.StatusInternalServerError)
		return
	}
	app := a.(apps.App)
	recordID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	result := utils.DB.Collection(app.Name).FindOne(ctx, bson.M{"_id": recordID})
	var record Record
	err = result.Decode(&record)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{"data": record})
}

// Create handler creates new log record
func Create(c *gin.Context) {
	a, exists := c.Get("app")
	if !exists {
		c.Status(http.StatusInternalServerError)
		return
	}
	app := a.(apps.App)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("failed to parse request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse request body"})
		return
	}

	var record Record

	err = json.Unmarshal(body, &record)
	if err != nil {
		log.Println("failed to unmarshal request body")
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to unmarshal request body"})
		return
	}

	record.ID = primitive.NewObjectIDFromTimestamp(time.Now().UTC())
	record.Created = time.Now().UTC()

	if app.Schema != "" {
		schemaLoader := gojsonschema.NewStringLoader(app.Schema)
		documentLoader := gojsonschema.NewBytesLoader(body)
		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			log.Println(err)
		}
		if !result.Valid() {
			var errors []string
			for _, desc := range result.Errors() {
				errors = append(errors, fmt.Sprintf("[%s] -> %s", desc.Field(), desc.Description()))
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("schema validation failed:\n%s", strings.Join(errors, "\n"))})
			return
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = utils.DB.Collection(app.Name).InsertOne(ctx, record)
	if err != nil {
		log.Println("failed to insert record in database")
	}
	c.JSON(http.StatusCreated, gin.H{"id": record.ID})

}

// DeleteAll handler clears app's logs
func DeleteAll(c *gin.Context) {
	a, exists := c.Get("app")
	if !exists {
		c.Status(http.StatusInternalServerError)
		return
	}
	app := a.(apps.App)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	utils.DB.Collection(app.Name).Drop(ctx)
}

// Delete handler deletes log by id
func Delete(c *gin.Context) {
	a, exists := c.Get("app")
	if !exists {
		c.Status(http.StatusInternalServerError)
		return
	}
	app := a.(apps.App)
	recordID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = utils.DB.Collection(app.Name).DeleteOne(ctx, bson.M{"_id": recordID})
	if err != nil {
		log.Println(err)
	}
	c.Status(http.StatusOK)
}

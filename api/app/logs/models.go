package logs

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Record struct {
	ID      primitive.ObjectID     `bson:"_id" json:"id,omitempty"`
	Data    map[string]interface{} `json:"data" binding:"required"`
	Created time.Time              `json:"created"`
}

package apps

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// App struct
type App struct {
	ID      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name    string             `json:"name" binding:"required"`
	Token   string             `json:"token,omitempty"`
	Schema  string             `json:"schema,omitempty"`
	Created time.Time          `json:"created,omitempty"`
}

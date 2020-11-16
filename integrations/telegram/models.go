package telegram

import "go.mongodb.org/mongo-driver/bson/primitive"

type Settings struct {
	Name  string `json:"name"`
	Token string `json:"token" binding:"required"`
	Users []User `json:"users"`
}

type User struct {
	Username string `json:"username"`
	Chat     int    `json:"chat"`
	Apps     []struct {
		ID    primitive.ObjectID `json:"id"`
		Level string             `json:"level"`
	} `json:"apps"`
}

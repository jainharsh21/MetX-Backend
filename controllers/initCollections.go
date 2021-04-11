package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func InitCollections(c *mongo.Database) {
	userCollection = c.Collection("users")
}

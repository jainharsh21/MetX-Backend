package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func GetUsers(c *gin.Context) {
	users := []User{}
	cursor, err := userCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var user User
		cursor.Decode(&user)
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Users",
		"data":    users,
	})
	return
}

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	name := user.Name
	id := guuid.New().String()

	newUser := User{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
	}

	_, err := userCollection.InsertOne(context.TODO(), newUser)

	if err != nil {
		log.Printf("Error while inserting new user into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created Successfully",
	})
	return
}

func GetUser(c *gin.Context) {
	userId := c.Param("userId")

	user := User{}
	err := userCollection.FindOne(context.TODO(), bson.M{"id": userId}).Decode(&user)
	if err != nil {
		log.Printf("Error while getting a single user, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single User",
		"data":    user,
	})
	return
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("userId")
	var user User
	c.BindJSON(&user)
	name := user.Name

	newData := bson.M{
		"$set": bson.M{
			"name": name,
		},
	}

	_, err := userCollection.UpdateOne(context.TODO(), bson.M{"id": userId}, newData)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "User Updated Successfully",
	})
	return
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")

	_, err := userCollection.DeleteOne(context.TODO(), bson.M{"id": userId})
	if err != nil {
		log.Printf("Error while deleting a single user, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User deleted successfully",
	})
	return
}
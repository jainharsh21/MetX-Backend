package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/jainharsh21/MetX-Backend/utils"
	"github.com/jainharsh21/MetX-Backend/models"
)



func GetUsers(c *gin.Context) {
	users := []models.User{}
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
		var user models.User
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
	var user models.User
	c.BindJSON(&user)
	id := guuid.New().String()
	name := user.Name
	email := user.Email
	phone := user.Phone
	password := user.Password
	user_type := user.UserType
	img_url := user.ImgUrl

	newUser := models.User{
		ID:        id,
		Name:      name,
		Email:     email,
		Phone:     phone,
		Password:  utils.HashPassword(password),
		UserType:  user_type,
		ImgUrl:    img_url,
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

	user := models.User{}
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
	var user models.User
	c.BindJSON(&user)
	name := user.Name
	phone := user.Phone
	img_url := user.ImgUrl
	log.Printf(img_url)

	newData := bson.M{
		"$set": bson.M{
			"name":    name,
			"phone":   phone,
			"imgurl": img_url,
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

func UserLogin(c *gin.Context) {

}

package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
	"github.com/jainharsh21/MetX-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

func GetEvents(c *gin.Context) {
	events := []models.Event{}
	cursor, err := eventCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all events, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var event models.Event
		cursor.Decode(&event)
		events = append(events, event)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Events",
		"data":    events,
	})
	return
}

func CreateEvent(c *gin.Context) {
	var event models.Event
	c.BindJSON(&event)
	id := guuid.New().String()
	name := event.Name
	img_url := event.ImgUrl
	description := event.Description
	summary := event.Summary
	event_at, _ := time.Parse(time.RFC3339, event.EventAt)
	location := event.Location
	fees := event.Fees
	student_chapter_id := event.StudentChapID
	tags := event.Tags

	fmt.Println(event_at)

	newEvent := models.Event{
		ID:            id,
		Name:          name,
		ImgUrl:        img_url,
		Description:   description,
		Summary:       summary,
		Location:      location,
		Fees:          fees,
		StudentChapID: student_chapter_id,
		Attendees:     []string{},
		Tags:          tags,
		// EventAt:       event.EventAt,
		EventAtOg: event_at,
	}

	_, err := eventCollection.InsertOne(context.TODO(), newEvent)

	if err != nil {
		log.Printf("Error while inserting new event into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Event created Successfully",
	})
	return
}

func GetEvent(c *gin.Context) {
	eventId := c.Param("eventId")

	event := models.Event{}
	err := eventCollection.FindOne(context.TODO(), bson.M{"id": eventId}).Decode(&eventId)
	if err != nil {
		log.Printf("Error while getting a single event, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Event not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Event",
		"data":    event,
	})
	return
}

func UpdateEvent(c *gin.Context) {
	eventId := c.Param("eventId")
	var event models.Event
	c.BindJSON(&event)
	name := event.Name
	img_url := event.ImgUrl
	description := event.Description
	summary := event.Summary
	event_at, _ := time.Parse(time.RFC3339, event.EventAt)
	location := event.Location
	fees := event.Fees
	tags := event.Tags

	newData := bson.M{
		"$set": bson.M{
			"name":        name,
			"imgurl":      img_url,
			"description": description,
			"summary":     summary,
			"eventatog":   event_at,
			"attendees":   []string{},
			"location":    location,
			"fees":        fees,
			"tags":        tags,
		},
	}

	_, err := eventCollection.UpdateOne(context.TODO(), bson.M{"id": eventId}, newData)
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
		"message": "Event Updated Successfully",
	})
	return
}

func DeleteEvent(c *gin.Context) {
	eventId := c.Param("eventId")

	_, err := eventCollection.DeleteOne(context.TODO(), bson.M{"id": eventId})
	if err != nil {
		log.Printf("Error while deleting a single event, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Event deleted successfully",
	})
	return
}

func AddAttendee(c *gin.Context) {
	eventId := c.Param("eventId")
	attendeeId := c.Param("attendeeId")
	newData := bson.M{
		"$push": bson.M{
			"attendees": attendeeId,
		},
	}

	_, err := eventCollection.UpdateOne(context.TODO(), bson.M{"id": eventId}, newData)
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
		"message": "Attendee Added Successfully",
	})
	return
}

func getAttendees() {

}

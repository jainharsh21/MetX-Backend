package config

import (
	"context"
	"github.com/jainharsh21/MetX-Backend/controllers"
	"github.com/jainharsh21/MetX-Backend/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func Connect() {
	// Database Config
	clientOptions := options.Client().ApplyURI(env.GetEnvVar("MONGO_URI"))
	client, err := mongo.NewClient(clientOptions)
	//Set up a context required by mongo.Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	//Cancel context to avoid memory leak
	defer cancel()

	// Ping our db connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}
	// Connect to the database
	db := client.Database("metx")
	controllers.InitCollections(db)
	return
}

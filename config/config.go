package config

import (
	"context"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/jainharsh21/MetX-Backend/env"
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
	client.Database("metx")
	return
	}
package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

type event struct {
	Title    string
	Time     int
	Price    int
	Genre    string
	Category []string
	Image    string
	Tickets  string
	Text     string
	Link     string
	Other    []string
	Location location
}

type location struct {
	Address     address
	Area        string
	Place       string
	Coordinates coordinates
}

type coordinates struct {
	Latitude  float64
	Longitude float64
}

type address struct {
	Street string
	No     string
	Zip    int
	City   string
	State  string
}

func getEvents(filter bson.M) (events []event) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := client.Database("dtu").Collection("moro").Find(ctx, filter)
	if err != nil {
		fmt.Println(err)
	}
	var event event
	for cursor.TryNext(context.Background()) {
		cursor.Decode(&event)
		events = append(events, event)
	}
	return events
}

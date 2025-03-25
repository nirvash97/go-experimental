package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb+srv://aegisx1:papth0391@experimental-01.8lsgx.mongodb.net/?retryWrites=true&w=majority&appName=Experimental-01"

func main() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	var echo bson.M
	if err := client.Database("admin").RunCommand(context.Background(), bson.D{{Key: "ping", Value: 1}}).Decode(&echo); err != nil {
		panic(err)
	}
	db := client.Database("sample_mflix")
	filter := bson.D{{Key: "languages", Value: "Thai"}}
	collection := db.Collection("movies")
	result, err := collection.Find(context.Background(), filter)

	for result.Next(context.Background()) {
		var movie Movie
		err := result.Decode(&movie)
		if err != nil {
			log.Fatal(err)

		}
		fmt.Printf("Title: %s\n", movie.Title)
		fmt.Printf("Runtime: %d minutes\n", movie.Runtime)
		fmt.Printf("Languages: %v\n", movie.Languages)
		fmt.Printf("Plot: %s\n\n", movie.Plot)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

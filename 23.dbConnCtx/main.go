package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	Task string `bson:"task"`
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))

	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	defer client.Disconnect(ctx)

	fmt.Println("Connected to MongoDB")

	collection := client.Database("tododb").Collection("todos")

	todo := Todo{Task: "learn go"}

	insertResult, err := collection.InsertOne(ctx, todo)

	if err != nil {
		fmt.Println("Insert error:", err)
		return
	}

	fmt.Println("Insert TODO with ID:", insertResult.InsertedID)

	cursor, err := collection.Find(ctx, map[string]interface{}{})

	if err != nil {
		fmt.Println("Find error:", err)
		return
	}

	defer cursor.Close(ctx)

	fmt.Println("TODO list:")

	for cursor.Next(ctx) {
		var t Todo
		cursor.Decode(&t)
		fmt.Println("-", t.Task)
	}

}

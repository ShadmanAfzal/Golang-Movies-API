package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "golang"
const collectionName = "movies"

var Collection *mongo.Collection

func init() {

	databaseURL := "mongodb+srv://Shadman:YN8MvnBruubF8zrN@cluster0.c4g4i.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

	collectionOption := options.Client().ApplyURI(databaseURL)

	client, err := mongo.Connect(context.TODO(), collectionOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("database connected successfully...")

	Collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("database instance is ready...")
}

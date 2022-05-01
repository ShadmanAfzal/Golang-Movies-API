package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/shadmanAfzal/movies_api/database"
	"github.com/shadmanAfzal/movies_api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertMovie(movie model.Movie) {
	inserted, err := database.Collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of record inserted", inserted)
}

func GetMovies() []primitive.M {
	cursor, err := database.Collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {

		var movie bson.M

		err = cursor.Decode(&movie)

		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())

	return movies
}

func DeleteMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	result, err := database.Collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total records deleted", result.DeletedCount)
}

func GetMovie(movieId string) primitive.M {
	id, _ := primitive.ObjectIDFromHex(movieId)

	var result bson.M

	filter := bson.M{"_id": id}

	err := database.Collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		return result
	}

	return result
}

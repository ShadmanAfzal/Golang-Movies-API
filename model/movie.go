package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	Id                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title              string             `json:"title,omitempty"`
	Cast               []Person           `json:"cast,omitempty"`
	Producer           Person             `json:"producer,omitempty"`
	Director           Person             `json:"director,omitempty"`
	IMDBRating         string             `json:"imdbRating,omitempty"`
	RottenTomatoRating string             `json:"rottenTomatoRating,omitempty"`
	MoviePoster        string             `json:"moviePoster,omitempty"`
	Trailer            string             `json:"trailer,omitempty"`
	Genre              []string           `json:"genre,omitempty"`
}

type Person struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	ProfileImage string `json:"profile_image,omitempty"`
}

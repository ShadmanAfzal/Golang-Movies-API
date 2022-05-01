package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shadmanAfzal/movies_api/helper"
	"github.com/shadmanAfzal/movies_api/model"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("api is working..."))
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	movies := helper.GetMovies()

	if len(movies) == 0 {
		w.Write([]byte("{}"))
	} else {
		json.NewEncoder(w).Encode(movies)
	}
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("{'status': 'json body can't be empty'}")
	}

	var movie model.Movie

	err := json.NewDecoder(r.Body).Decode(&movie)

	if err != nil {
		log.Fatal(err)
		w.Write([]byte("{'status': 'invalid json value'}"))
		return
	}

	helper.InsertMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	helper.DeleteMovie(params["id"])

	json.NewEncoder(w).Encode("{'status': 'movie deleted successfully'}")
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	result := helper.GetMovie(params["id"])

	if result == nil {
		w.Write([]byte("{}"))
		return
	}

	json.NewEncoder(w).Encode(result)
}

package router

import (
	"github.com/gorilla/mux"
	"github.com/shadmanAfzal/movies_api/controller"
)

func GetRouter() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", controller.Home)
	route.HandleFunc("/api/v1/movies", controller.GetMovies).Methods("GET")
	route.HandleFunc("/api/v1/movie", controller.AddMovie).Methods("POST")
	route.HandleFunc("/api/v1/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	route.HandleFunc("/api/v1/movie/{id}", controller.GetMovie).Methods("GET")
	return route
}

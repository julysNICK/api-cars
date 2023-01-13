package carsController

import (
	"apicars/middlewares"
	"net/http"

	"github.com/go-chi/cors"
)

func (r *ServerConfig) GetRouter() {

	r.Router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)
	r.Router.HandleFunc("/api/v1/", HelloApi).Methods("GET")
	r.Router.HandleFunc("/api/v1/car/{id}", middlewares.SetMiddlewareAuthentication(r.GetCarById)).Methods("GET")
	r.Router.HandleFunc("/api/v1/car", middlewares.SetMiddlewareAuthentication(r.AddCar)).Methods("POST")
	r.Router.HandleFunc("/api/v1/cars/my" , middlewares.SetMiddlewareAuthentication(r.GetCarsByMyIdUser)).Methods("GET")
	r.Router.HandleFunc("/api/v1/cars/user/{id}" , middlewares.SetMiddlewareAuthentication(r.GetCarsByUserId)).Methods("GET")
	r.Router.HandleFunc("/api/v1/cars", r.GetCars).Methods("GET")
	r.Router.HandleFunc("/api/v1/car/{id}", r.UpdateCar).Methods("PATCH")
	r.Router.HandleFunc("/api/v1/car/{id}", r.DeleteCar).Methods("DELETE")

	r.Router.HandleFunc("/api/v1/register", r.Register).Methods("POST", http.MethodOptions)
	r.Router.HandleFunc("/api/v1/login", r.Login).Methods("POST",http.MethodOptions)
	r.Router.HandleFunc("/api/v1/refresh", r.RefreshSession).Methods("POST")
}

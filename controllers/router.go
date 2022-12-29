package carsController

import "apicars/middlewares"

func (r *ServerConfig) GetRouter() {
	r.Router.HandleFunc("/api/v1/", HelloApi).Methods("GET")
	r.Router.HandleFunc("/api/v1/car/{id}", middlewares.SetMiddlewareAuthentication(r.GetCarById)).Methods("GET")
	r.Router.HandleFunc("/api/v1/car", middlewares.SetMiddlewareAuthentication(r.AddCar)).Methods("POST")
	r.Router.HandleFunc("/api/v1/cars", r.GetCars).Methods("GET")
	r.Router.HandleFunc("/api/v1/car/{id}", r.UpdateCar).Methods("PATCH")
	r.Router.HandleFunc("/api/v1/car/{id}", r.DeleteCar).Methods("DELETE")

	r.Router.HandleFunc("/api/v1/register", r.Register).Methods("POST")
	r.Router.HandleFunc("/api/v1/login", r.Login).Methods("POST")
	r.Router.HandleFunc("/api/v1/refresh", r.RefreshSession).Methods("POST")
}

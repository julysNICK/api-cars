package carsController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"apicars/auth"
	"apicars/models"
	"apicars/services"

	"github.com/gorilla/mux"
)

var CarList = []models.Car{
	{Make: "Ford", Model: "Mustang", Year: "1969", Is_Sold: false},
	{Make: "Ford", Model: "F150", Year: "2018", Is_Sold: true},
	{Make: "Chevrolet", Model: "Camaro", Year: "2019", Is_Sold: false},
	{Make: "Chevrolet", Model: "Silverado", Year: "2018", Is_Sold: true},
	{Make: "Dodge", Model: "Charger", Year: "2019", Is_Sold: false},
}

func HelloApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello, API"}`))
}

func (ServerConfig *ServerConfig) GetCars(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	carsList, err := services.GetAllCars(ServerConfig.DB)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Cars not found"}`))
		return
	}

	json.NewEncoder(w).Encode(carsList)

}

func (ServerConfig *ServerConfig) GetCarById(w http.ResponseWriter, r *http.Request) {
	idCar := mux.Vars(r)["id"]
	if idCar == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Car not found - id not found"}`))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	car, err := services.GetCarById(ServerConfig.DB, idCar)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Car not found"}`))
		return
	}

	json.NewEncoder(w).Encode(car)
}

func (ServerConfig *ServerConfig) AddCar(w http.ResponseWriter, r *http.Request) {
	var newCar models.Car
	_ = json.NewDecoder(r.Body).Decode(&newCar)

	uid, err := auth.ExtractTokenId(r)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Unauthorized"}`))
		return
	}
	errCar := services.CreateCar(ServerConfig.DB, newCar, uid)

	if errCar != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Car not created"}`))
	}

	fmt.Println(uid)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Car created"}`))
}

func (ServerConfig *ServerConfig) UpdateCar(w http.ResponseWriter, r *http.Request) {
	idCar := mux.Vars(r)["id"]
	if idCar == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Car not updated - id not found"}`))
	}

	var newCar models.Car
	_ = json.NewDecoder(r.Body).Decode(&newCar)

	err := services.UpdateCarById(ServerConfig.DB, idCar, newCar)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Car not updated"}`))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Car updated"}`))

}

func (ServerConfig *ServerConfig) DeleteCar(w http.ResponseWriter, r *http.Request) {
	idCar := mux.Vars(r)["id"]

	if idCar == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Car not deleted - id not found"}`))
	}

	err := services.DeleteCarById(ServerConfig.DB, idCar)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Car not deleted"}`))
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Car deleted"}`))

}

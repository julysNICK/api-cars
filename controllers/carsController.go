package carsController

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"apicars/auth"
	"apicars/models"
	"apicars/services"

	utilsResponse "apicars/utils"

	"github.com/gorilla/mux"
)

func HelloApi(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello, API"}`))
}

func (ServerConfig *ServerConfig) GetCars(w http.ResponseWriter, r *http.Request) {

	carsList, err := services.GetAllCars(ServerConfig.DB)

	if err != nil {
		utilsResponse.ResponseError(w, http.StatusNotFound, "Cars not found")
		return
	}

	utilsResponse.ResponseJson(w, http.StatusOK, carsList)

}

func (ServerConfig *ServerConfig) GetCarById(w http.ResponseWriter, r *http.Request) {
	idCar := mux.Vars(r)["id"]
	if idCar == "" {
		utilsResponse.ResponseError(w, http.StatusBadRequest, "Id is required")
		return
	}

	car, err := services.GetCarById(ServerConfig.DB, idCar)

	if err != nil {
		utilsResponse.ResponseError(w, http.StatusNotFound, "Car not found")
		return
	}

	utilsResponse.ResponseJson(w, http.StatusOK, car)
}

func (ServerConfig *ServerConfig) GetCarsByMyIdUser(w http.ResponseWriter, r *http.Request) {
	uid, err := auth.ExtractTokenId(r)
	if err != nil {
		utilsResponse.ResponseError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	userCars, err, err2 := services.GetCarsByMyIdUser(ServerConfig.DB, uid)

	if err != nil || err2 != nil {
		utilsResponse.ResponseError(w, http.StatusNotFound, "Cars not found")
		return
	}

	utilsResponse.ResponseJson(w, http.StatusOK, userCars)
}

func (ServerConfig *ServerConfig) GetCarsByUserId(w http.ResponseWriter, r *http.Request) {
	uid := mux.Vars(r)["id"]
	convertInt, _ := strconv.Atoi(uid)

	if uid == "" {
		utilsResponse.ResponseError(w, http.StatusBadRequest, "Id is required")
		return
	}

	_, err, err2 := services.GetCarsByMyIdUser(ServerConfig.DB, uint(convertInt))

	if err != nil || err2 != nil {
		utilsResponse.ResponseError(w, http.StatusNotFound, "Cars not found")
		return
	}

	utilsResponse.ResponseJson(w, http.StatusOK, "oi")
}

func (ServerConfig *ServerConfig) AddCar(w http.ResponseWriter, r *http.Request) {
	var newCar models.Car
	_ = json.NewDecoder(r.Body).Decode(&newCar)

	uid, err := auth.ExtractTokenId(r)

	if err != nil {
		utilsResponse.ResponseError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	errCar := services.CreateCar(ServerConfig.DB, newCar, uid)

	if errCar != nil {
		utilsResponse.ResponseError(w, http.StatusInternalServerError, "Error creating car")
		return
	}

	utilsResponse.ResponseJson(w, http.StatusCreated, newCar)
}

func (ServerConfig *ServerConfig) UpdateCar(w http.ResponseWriter, r *http.Request) {
	idCar := mux.Vars(r)["id"]
	if idCar == "" {
		utilsResponse.ResponseError(w, http.StatusBadRequest, "Id is required")
	}

	var newCar models.Car
	_ = json.NewDecoder(r.Body).Decode(&newCar)

	err := services.UpdateCarById(ServerConfig.DB, idCar, newCar)

	if err != nil {
		utilsResponse.ResponseError(w, http.StatusNotFound, "Car not updated")
		return
	}
	utilsResponse.ResponseJson(w, http.StatusOK, newCar)
}

func (ServerConfig *ServerConfig) DeleteCar(w http.ResponseWriter, r *http.Request) {
	idCar := mux.Vars(r)["id"]

	if idCar == "" {
		utilsResponse.ResponseError(w, http.StatusBadRequest, "Id is required")
		return
	}

	err := services.DeleteCarById(ServerConfig.DB, idCar)

	if err != nil {
		utilsResponse.ResponseError(w, http.StatusNotFound, "Car not deleted")
		return
	}

	utilsResponse.ResponseJson(w, http.StatusOK, fmt.Sprintf("Car %s deleted", idCar))

}

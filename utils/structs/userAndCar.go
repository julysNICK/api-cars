package structs

import "apicars/models"

type UserInfo struct {
	Id        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type CarsUser struct {
	Cars []models.Car
	User models.User
}

type CarsUserUnique struct {
	Car  models.Car
	User UserInfo
}

type ListCarsUser struct {
	CarsInfo CarsUserUnique
}

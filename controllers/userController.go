package carsController

import (
	"apicars/models"
	"apicars/services"
	"apicars/utils/structs"
	"fmt"

	utilsResponse "apicars/utils"
	"encoding/json"
	"net/http"
)


func (ServerConfig *ServerConfig) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	user, err := services.CreateUser(ServerConfig.DB, user)
	if err != nil {
		utilsResponse.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	utilsResponse.ResponseJson(w, http.StatusOK, structs.ResponseUser{Message: "User created successfully", User: user})
}

func (ServerConfig *ServerConfig) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	userFound, token, err := services.Login(ServerConfig.DB, user.Email, user.Password)

	if err != "" {
		utilsResponse.ResponseError(w, http.StatusUnauthorized, err)
		return
	}
	utilsResponse.ResponseJson(w, http.StatusOK, structs.ResponseUser{Message: "User logged successfully", User: userFound, Token: token})
}

func (ServerConfig *ServerConfig) RefreshSession(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	newToken, err := services.RefreshToken(ServerConfig.DB, r)

	if err != nil {
		utilsResponse.ResponseError(w, http.StatusUnauthorized, err.Error())
		return
	}

	utilsResponse.ResponseJson(w, http.StatusOK, structs.RefreshTokenResponse{Message: "Token refreshed successfully", Token: newToken})

}

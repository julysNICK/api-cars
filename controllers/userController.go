package carsController

import (
	"apicars/models"
	"apicars/services"
	"apicars/structs"
	"encoding/json"
	"net/http"
)


func (ServerConfig *ServerConfig) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user, err := services.CreateUser(ServerConfig.DB, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message": "Error creating user"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode( structs.ResponseUser{Message: "User created successfully", User: user})
}

func (ServerConfig *ServerConfig) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	userFound, token, err := services.Login(ServerConfig.DB, user.Email, user.Password)
	if err != "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Error login user"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(structs.ResponseUser{Message: "User login successfully", User: userFound, Token: token})
}

func (ServerConfig *ServerConfig) RefreshSession(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	newToken, err := services.RefreshToken(ServerConfig.DB, r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Error refreshing token"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(structs.ResponseUser{Message: "Token refreshed successfully", Token: newToken})

}

package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, status int, data interface{}){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func ResponseError(w http.ResponseWriter, status int, err string){
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}
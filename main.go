package main

import (
	BdConfig "apicars/controllers"
	internalBdConfig "apicars/internal"

	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	serverConfig := BdConfig.ServerConfig{}
	conn, router := serverConfig.Inicialize()
	serverConfig.GetRouter()

	for _, seed := range internalBdConfig.All() {
		err := seed.Run(
			conn,
		)
		if err != nil {
			panic(err)
		}
	}

	http.ListenAndServe(":8080", router)

}

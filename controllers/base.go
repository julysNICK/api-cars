package carsController

import (
	"os"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	internal "apicars/internal"
)

type ServerConfig struct {
	// essa struct vai pode ser usada com o valor daqui em qualquer lugar
	DB     *gorm.DB
	Router *mux.Router
}

func (ServerConfig *ServerConfig) Inicialize() (*gorm.DB, *mux.Router) {
	ServerConfig.Router = mux.NewRouter()

	ServerConfig.DB = internal.GetConnection(
		os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("TEST_DB_HOST"), os.Getenv("DB_NAME"),
	)
	return ServerConfig.DB, ServerConfig.Router
}

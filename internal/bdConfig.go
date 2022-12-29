package bdConfig

import (
	"apicars/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	psqlMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *gorm.DB {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName, DbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening database-L20: ", err)
	}
	log.Println("Connection Opened to Database")
	db.AutoMigrate(&models.User{}, &Cars{})
	return db

}

func GetConnectionMigration() {

	log.Println("Connection Migration to Database")

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")),
	)

	if err != nil {
		log.Fatal("Error opening database-L36: ", err)
	}

	driver, err := psqlMigrate.WithInstance(db, &psqlMigrate.Config{})

	if err != nil {
		log.Fatal("Error opening database-L42: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migrations",
		"postgres", driver)

	if err != nil {
		log.Fatal("Error opening database-L50: ", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Error opening database-L54: ", err)
	}

	log.Println("Connection Migration to Database")
}

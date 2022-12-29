package services

import (
	"apicars/models"
	"log"
	"net/http"

	"apicars/auth"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUser(db *gorm.DB, id string) (models.User, error) {
	var user models.User
	err := db.Where("id = ?", id).First(&user).Error
	return user, err
}

func GetUsers(db *gorm.DB) ([]models.User, error) {
	var users []models.User
	err := db.Find(&users).Error
	return users, err
}

func CreateUser(db *gorm.DB, user models.User) (models.User, error) {
	hashedPassword, _ := hashPassword(user.Password)
	user.Password = string(hashedPassword)
	err := db.Create(&user).Error
	return user, err
}

func GetUserByEmail(db *gorm.DB, email string) (models.User, error) {
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	return user, err
}

func Login(db *gorm.DB, email, password string) (models.User, string, string) {
	user, err := GetUserByEmail(db, email)

	if err != nil {
		log.Println("ErrorL53: ", err)
		return models.User{}, "", "User not found"
	}
	if !CheckPasswordHash(password, user.Password) {

		return models.User{}, "", "Invalid password"
	}

	token, err := auth.CreateToken(user.ID)

	if err != nil {

		return models.User{}, "", "Error creating token"
	}

	return user, token, ""

}

func RefreshToken(db *gorm.DB, r *http.Request) (string, error) {
	tokenExpired := auth.ExtractToken(r)
	tokenExpiredId, err := auth.ExtractTokenExpiredId(r)

	if err != nil {
		return "", err
	}

	if tokenExpired == "" {

		return "", nil
	}

	newToken, err := auth.RefreshToken(tokenExpired, tokenExpiredId)

	if err != nil {
		return "", err
	}

	return newToken, nil

}

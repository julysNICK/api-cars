package services

import (
	"apicars/models"

	"gorm.io/gorm"
)

func GetAllCars(db *gorm.DB) ([]models.Car, error) {
	var cars []models.Car
	err := db.Find(&cars).Error
	return cars, err
}

func GetCarById(db *gorm.DB, id string) (models.Car, error) {
	var car models.Car
	err := db.Where("id = ?", id).First(&car).Error
	return car, err
}

func GetCarByModel(db *gorm.DB, model string) (models.Car, error) {
	var car models.Car
	err := db.Where("model = ?", model).First(&car).Error
	return car, err
}

func GetCarsByYear(db *gorm.DB, year int) ([]models.Car, error) {
	var cars []models.Car
	err := db.Where("year = ?", year).Find(&cars).Error
	return cars, err
}

func GetCarsByMake(db *gorm.DB, make string) ([]models.Car, error) {
	var cars []models.Car
	err := db.Where("make = ?", make).Find(&cars).Error
	return cars, err
}

func GetCarsBySold(db *gorm.DB, sold bool) ([]models.Car, error) {
	var cars []models.Car
	err := db.Where("is_sold = ?", sold).Find(&cars).Error
	return cars, err
}

func CreateCar(db *gorm.DB, car models.Car, user_id uint) error {
	car.User_Id = user_id
	err := db.Create(&car).Error
	return err
}

func UpdateCar(db *gorm.DB, car models.Car) error {
	err := db.Save(&car).Error
	return err
}

func UpdateCarById(db *gorm.DB, id string, car models.Car) error {
	err := db.Model(&car).Where("id = ?", id).Updates(car).Error
	return err
}

func DeleteCar(db *gorm.DB, car models.Car) error {
	err := db.Delete(&car).Error
	return err
}

func DeleteCarById(db *gorm.DB, id string) error {
	err := db.Where("id = ?", id).Delete(&models.Car{}).Error
	return err
}

func DeleteAllCars(db *gorm.DB) error {
	err := db.Unscoped().Delete(&models.Car{}).Error
	return err
}

func GetCarsByYearAndMake(db *gorm.DB, year int, make string) ([]models.Car, error) {
	var cars []models.Car
	err := db.Where("year = ? AND make = ?", year, make).Find(&cars).Error
	return cars, err
}

func GetCarsByYearAndSold(db *gorm.DB, year int, sold bool) ([]models.Car, error) {
	var cars []models.Car
	err := db.Where("year = ? AND is_sold = ?", year, sold).Find(&cars).Error
	return cars, err
}

func GetCarsByMakeAndSold(db *gorm.DB, make string, sold bool) ([]models.Car, error) {
	var cars []models.Car
	err := db.Where("make = ? AND is_sold = ?", make, sold).Find(&cars).Error
	return cars, err
}

func GetCarsByYearAndMakeAndSold(db *gorm.DB, year int, make string, sold bool) ([]models.Car, error) {
	var cars []models.Car
	err := db.Where("year = ? AND make = ? AND is_sold = ?", year, make, sold).Find(&cars).Error
	return cars, err
}

func GetCarsByYearAndMakeAndSoldAndModel(db *gorm.DB, year int, make string, sold bool, model string) ([]models.Car, error) {
	var cars []models.Car
	err := db.Where("year = ? AND make = ? AND is_sold = ? AND model = ?", year, make, sold, model).Find(&cars).Error
	return cars, err
}

func GetCarsByYearOrMake(db *gorm.DB, year int, make string) ([]models.Car, error) {
	var cars []models.Car
	err := db.Where("year = ? OR make = ?", year, make).Find(&cars).Error
	return cars, err
}

func GetCarsByYearOrSold(db *gorm.DB, year int, sold bool) ([]models.Car, error) {
	var cars []models.Car
	err := db.Where("year = ? OR is_sold = ?", year, sold).Find(&cars).Error
	return cars, err
}

func GetCarsByMakeOrSold(db *gorm.DB, make string, sold bool) ([]models.Car, error) {
	var cars []models.Car
	err := db.Where("make = ? OR is_sold = ?", make, sold).Find(&cars).Error
	return cars, err
}

func GetCarsByYearOrMakeOrSold(db *gorm.DB, year int, make string, sold bool) ([]models.Car, error) {
	var cars []models.Car
	err := db.Where("year = ? OR make = ? OR is_sold = ?", year, make, sold).Find(&cars).Error
	return cars, err
}

func GetCarsByYearOrMakeOrSoldOrModel(db *gorm.DB, year int, make string, sold bool, model string) ([]models.Car, error) {
	var cars []models.Car
	err := db.Where("year = ? OR make = ? OR is_sold = ? OR model = ?", year, make, sold, model).Find(&cars).Error
	return cars, err
}
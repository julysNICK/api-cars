package bdConfig

import (
	"apicars/models"
	"time"

	"gorm.io/gorm"
)

type Cars struct {
	Make            string    `json:"make"`
	Description_Car string    `json:"descriptionCar"`
	Image           string    `json:"image"`
	Model           string    `json:"model"`
	Reason_To_Sell  string    `json:"reasonToSell"`
	Year            string    `json:"year"`
	Is_Sold         bool      `json:"isSold"`
	Color           string    `json:"color"`
	Price           int       `json:"price"`
	User_Id         uint      `json:"user_id"`
	CreatedAt       time.Time `json:"-"`
	UpdatedAt       time.Time `json:"-"`
}

type Seed struct {
	Cars Cars
	Run  func(*gorm.DB) error
}

func CreateCar(db *gorm.DB, car Cars) error {
	return db.Create(&car).Error
}

func All() []Seed {
	return []Seed{
		{
			Cars: Cars{
				Model:          "Ferrari",
				Make:           "Ferrari",
				Image:          "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSZrounUeNxS-_8IzjiG_PYuqjDdv0J5uk6Zw&usqp=CAU",
				Year:           "2020",
				Reason_To_Sell: "I want to buy a new car",
				Color:          "Red",
				Price:          1000000,
				Is_Sold:        false,
				User_Id:        1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			Run: func(d *gorm.DB) error {
				return CreateCar(d, Cars{
					Model:          "Ferrari",
					Make:           "Ferrari",
					Image:          "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSZrounUeNxS-_8IzjiG_PYuqjDdv0J5uk6Zw&usqp=CAU",
					Year:           "2020",
					Reason_To_Sell: "I want to buy a new car",
					Color:          "Red",
					Price:          1000000,
					Is_Sold:        false,
					User_Id:        1,
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				})
			},
		},
		{
			Cars: Cars{
				Model:          "Ford",
				Make:           "Ford",
				Image:          "https://pictures.dealer.com/o/offleaseonlycommiamihomedeliverynow/1462/3a743def55179eced6185b57d00cc35bx.jpg?impolicy=downsize_bkpt&imdensity=1&w=520",
				Year:           "2002",
				Reason_To_Sell: "I want to buy a new car",
				Color:          "blue",
				Price:          1000000,
				Is_Sold:        false,
				User_Id:        1,
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			Run: func(d *gorm.DB) error {
				return CreateCar(d, Cars{
					Model:          "Ford",
					Make:           "Ford",
					Image:          "https://pictures.dealer.com/o/offleaseonlycommiamihomedeliverynow/1462/3a743def55179eced6185b57d00cc35bx.jpg?impolicy=downsize_bkpt&imdensity=1&w=520",
					Year:           "2002",
					Reason_To_Sell: "I want to buy a new car",
					Color:          "blue",
					Price:          1000000,
					Is_Sold:        false,
					User_Id:        1,
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				})
			},
		},
	}
}

type seedUsers struct {
	User models.User
	Run  func(*gorm.DB) error
}

func CreateUser(db *gorm.DB, user models.User) error {
	return db.Create(&user).Error
}

func AllUsers() []seedUsers {
	return []seedUsers{
		{
			User: models.User{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "test@gmail.com",
				Password:  "123456",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Run: func(d *gorm.DB) error {
				return CreateUser(d, models.User{
					FirstName: "John",
					LastName:  "Doe",
					Email:     "test@gmail.com",
					Password:  "123456",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				})
			},
		},
	}
}

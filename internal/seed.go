package bdConfig

import (
	"apicars/models"
	"time"

	"gorm.io/gorm"
)

type Cars struct {
	Id        int       `json:"id"`
	Model     string    `json:"model"`
	Make      string    `json:"make"`
	Year      string    `json:"year"`
	Color     string    `json:"color"`
	Price     int       `json:"price"`
	Is_Sold   bool      `json:"Is_Sold"`
	User_Id   int       `json:"user_id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
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
				Model:     "Ferrari",
				Make:      "Ferrari",
				Year:      "2020",
				Color:     "Red",
				Price:     1000000,
				Is_Sold:   false,
				User_Id:   1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Run: func(d *gorm.DB) error {
				return CreateCar(d, Cars{
					Model:     "Ferrari",
					Make:      "Ferrari",
					Year:      "2020",
					Color:     "Red",
					Price:     1000000,
					Is_Sold:   false,
					User_Id:   1,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				})
			},
		},
		{
			Cars: Cars{
				Model:     "Ford",
				Make:      "Ford",
				Year:      "2002",
				Color:     "blue",
				Price:     1000000,
				Is_Sold:   false,
				User_Id:   1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Run: func(d *gorm.DB) error {
				return CreateCar(d, Cars{
					Model:     "Ford",
					Make:      "Ford",
					Year:      "2002",
					Color:     "blue",
					Price:     1000000,
					Is_Sold:   false,
					User_Id:   1,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
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
				Password: "123456",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Run: func(d *gorm.DB) error {
				return CreateUser(d, models.User{
					FirstName: "John",
					LastName:  "Doe",
					Email:     "test@gmail.com",
					Password: "123456",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				})
			},
		},
	}
}
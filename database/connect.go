package database

import (
	"fmt"
	"strconv"

	"github.com/sohhamm/product-feedback-server/config"
	"github.com/sohhamm/product-feedback-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	p := config.Config("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println("Wrong port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("Failed to connect to the database 😥")
	}

	fmt.Println("Connected to the database 🔥")

	DB.AutoMigrate(&models.ProductRequest{})

}

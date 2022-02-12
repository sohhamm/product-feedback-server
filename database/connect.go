package database

import (
	"fmt"
	"log"
	"os"

	"github.com/sohhamm/product-feedback-server/config"
	"github.com/sohhamm/product-feedback-server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func ConnectDB() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s Timezone=Asia/Calcutta", config.Config("DB_HOST"), config.Config("DB_PORT"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), config.Config("SSL"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Error connecting to the db 😞 \n", err)
		os.Exit(2)
	}

	log.Println("Connected to the database 🔥")

	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations .... 🚂")

	db.AutoMigrate(&models.Feedback{}, &models.Comment{})

	DB = DbInstance{Db: db}

}

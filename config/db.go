package config

import (
	"fmt"
	"log"
	"os"

	"github.com/eshdc/content-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=eshdc_admin password=password dbname=eshdc_content port=5432 sslmode=disable"
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Content Database connected")

	// Migrate models
	DB.AutoMigrate(&models.News{}, &models.PageContent{}, &models.Media{}, &models.Setting{}, &models.HeroSlide{}, &models.Job{}, &models.Memo{}, &models.ContactMessage{})
}

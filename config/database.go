package config

import (
	"go-learn-restapi-mysql/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load variabel lingkungan dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Membaca variabel lingkungan untuk koneksi database
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Menghubungkan ke database MySQL
	database, err := gorm.Open(mysql.Open(dbUser + ":" + dbPassword + "@(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"))
	if err != nil {
		return
	}

	database.AutoMigrate(&models.Blog{}, &models.Users{})

	DB = database
}

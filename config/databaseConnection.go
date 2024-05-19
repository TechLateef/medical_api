package config

import (
	"fmt"
	"log"
	"medical_api/model"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env file:", err)
	}

	dbDatabase := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbDatabase == "" || dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" {
		log.Fatal("Missing required environment variables")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPassword, dbDatabase, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate all models
	err = db.AutoMigrate(&model.Patient{}, &model.Doctor{}, &model.MedicalRecord{})
	if err != nil {
		log.Fatal("Failed to perform auto migration:", err)
	}

	return db
}

func CloseDatabase(db *gorm.DB) {
	dbPostgre, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database connection:", err)
	}
	err = dbPostgre.Close()
	if err != nil {
		log.Fatal("Failed to close database connection:", err)
	}
}

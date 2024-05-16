package config

import (
	"fmt"
	"medical_api/model"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	dbDatabase := os.Getenv("DB_DATABASE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	fmt.Print(dbDatabase)
	fmt.Println(dbHost)
	fmt.Println(dbPort)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPassword, dbDatabase)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to load env file")

	}
	db.AutoMigrate(&model.User{})
	return db

}

func CloseDatabasec(db *gorm.DB) {
	dbPostgre, err := db.DB()
	if err != nil {
		panic("Failed ")
	}
	dbPostgre.Close()
}

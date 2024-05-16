package main

import (
	"medical_api/config"
	"medical_api/router"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
)

func main() {
	defer config.CloseDatabasec(db)

	router.Routes()

}

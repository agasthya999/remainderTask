package migrations

import (
	"fmt"
	"log"
	database "remainderTask/database"
	"remainderTask/models"
)

func Migrate() {
	database.GormConnect()

	err := database.DBRead.AutoMigrate(
		&models.User{},
		&models.Task{},
	)
	if err != nil {
		fmt.Println("Error while migrating ", err)
		return
	}
	log.Println("Succesfully migrated")
}

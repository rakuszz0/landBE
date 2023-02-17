package database

import (
	"fmt"
	"landtick/models"
	"landtick/pkg/connection"
)

func RunMigration() {
	err := connection.DB.AutoMigrate(
		&models.User{},
		&models.Penumpang{},
		&models.Stasiun{},
		&models.Train{},
		&models.Tiket{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Failed! Create Table to Database")
	}
	fmt.Println("Create Table Success")
}

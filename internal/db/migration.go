package db

import "hng11task2/internal/models"

func Migrate() error {
	err := DB.AutoMigrate(
		&models.Organisation{},
		&models.User{},
	)
	return err
}

package database

import (
	"gorm.io/gorm"
)

func ProvideDatabase() *gorm.DB {
	database, err := OpenDatabase()

	if err != nil {
		panic(err)
	}

	return database
}

package database

import (
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func OpenDatabase() (*gorm.DB, error) {
	if database != nil {
		return database, nil
	}

	path := viper.GetString("DATABASE")
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	database = db
	return database, nil
}

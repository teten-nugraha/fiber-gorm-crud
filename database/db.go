package database

import (
	"github.com/teten-nugraha/golang-crud-fiber/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() error {

	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&model.Product{})

	return nil
}
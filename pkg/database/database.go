package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	//dsn := "root:root@tcp(127.0.0.1:3306)/schedule?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "host=localhost user=kuppe_user password=UI8diYYhkkdm&&HG^dk dbname=kuppe_db port=15221 sslmode=disable"
	//dsn := "host=localhost user=postgres password=postgres dbname=kuppe_db port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB.AutoMigrate(&Lesson{})
	return nil
}

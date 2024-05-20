package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"src/ex02/models"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost user=postgres dbname=db password=postgres sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	//defer DB.Close()

	DB.AutoMigrate(&models.Post{})
}

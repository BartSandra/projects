package db

import (
	"log"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Record struct {
	//gorm.Model
	Uuid      string `gorm:"primaryKey"`
	Name      string
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func Connect(port int) (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=" + strconv.Itoa(port) + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к psql: %v", err)
	}

	if err = db.AutoMigrate(&Record{}); err != nil {
		return nil, err
	}
	return db, nil
}

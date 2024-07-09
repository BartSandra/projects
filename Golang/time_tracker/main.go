package main

import (
	"github.com/joho/godotenv"
	"time_tracker/api"
	"time_tracker/database"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel) 
}

func main() {
	log.Debug("Starting the application")

	err := godotenv.Load()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error loading .env file")
	}

	log.Debug("Loaded the .env file")

	err = database.Migrate()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error migrating the database")
	}

	log.Debug("Database migration completed")

	api.Run()

	log.Debug("Started the API server")
}


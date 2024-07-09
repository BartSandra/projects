package api

import (
	"net/http"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func Run() {
	log.Debug("Starting the server")
	log.Info("Listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", NewRouter()))
}


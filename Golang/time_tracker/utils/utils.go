package utils

import (
	"net/http"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	log.Debug("Responding with JSON") // Добавьте Debug логи
	response, err := json.Marshal(payload)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error marshalling JSON") // Добавьте Error логи
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func RespondError(w http.ResponseWriter, status int, message string) {
	log.Debug("Responding with error") // Добавьте Debug логи
	RespondJSON(w, status, map[string]string{"error": message})
}


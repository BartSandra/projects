package middleware

import (
	"net/http"
	"time"
	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		log.Debug("Handling request")

		next.ServeHTTP(w, r)

		duration := time.Now().Sub(startTime)
		log.WithFields(log.Fields{
			"method": r.Method,
			"uri": r.RequestURI,
			"duration": duration,
		}).Info("Handled request")

		log.Debug("Finished handling request")
	})
}

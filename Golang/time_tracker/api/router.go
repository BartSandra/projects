package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}/start", StartTask).Methods("POST")
	router.HandleFunc("/users/{id}/end", EndTask).Methods("POST")
	router.HandleFunc("/users/{id}/tasks", GetTasksByUserAndPeriod).Methods("GET")

	return router
}

package model

import "time"

type User struct {
	ID             string `json:"id"`
	PassportNumber string `json:"passportNumber"`
	PassportSerie  string `json:"passportSerie"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	Patronymic     string `json:"patronymic"`
	Address        string `json:"address"`
}

type Task struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

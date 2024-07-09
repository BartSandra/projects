Реализовать тайм-трекер на Go
Необходимо реализовать следующее
1. Выставить REST методы
Получение данных пользователей:
Фильтрация по всем полям.
Пагинация.
Получение трудозатрат по пользователю за период задача-сумма часов и минут с сортировкой от большей затраты к меньшей
Начать отсчет времени по задаче для пользователя
Закончить отсчет времени по задаче для пользователя
Удаление пользователя
Изменение данных пользователя
Добавление нового пользователя в формате:
```json
{
	"passportNumber": "1234 567890" // серия и номер паспорта пользователя
}
```
2. При добавлении сделать запрос в АПИ, описанного сваггером
```yaml
openapi: 3.0.3
info:
  title: People info
  version: 0.0.1
paths:
  /info:
    get:
      parameters:
        - name: passportSerie
          in: query
          required: true
          schema:
            type: integer
        - name: passportNumber
          in: query
          required: true
          schema:
            type: integer
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/People'
        '400':
          description: Bad request
        '500':
          description: Internal server error
components:
  schemas:
    People:
      required:
        - surname
        - name
        - address
      type: object
      properties:
        surname:
          type: string
          example: Иванов
        name:
          type: string
          example: Иван
        patronymic:
          type: string
          example: Иванович
        address:
          type: string
          example: г. Москва, ул. Ленина, д. 5, кв. 1
```
3. Обогащенную информацию положить в БД postgres (структура БД должна быть создана путем миграций при старте сервиса)
4. Покрыть код debug- и info-логами
5. Вынести конфигурационные данные в .env-файл
6. Сгенерировать сваггер на реализованное АПИ

Структура:

time_tracker/
│
├── main.go
│
├── api/
│   ├── handler.go
│   ├── router.go
│   └── server.go
│
├── config/
│   └── config.go
│
├── database/
│   └── db.go
│
├── middleware/
│   └── middleware.go
│
├── model/
│   └── user.go
│
└── utils/
    └── utils.go

Вот код:

time_tracker/main.go:

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

time_tracker/swagger.yaml:

openapi: 3.0.3
info:
  title: People info
  version: 0.0.1
paths:
  /info:
    get:
      parameters:
        - name: passportSerie
          in: query
          required: true
          schema:
            type: integer
        - name: passportNumber
          in: query
          required: true
          schema:
            type: integer
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/People'
        '400':
          description: Bad request
        '500':
          description: Internal server error
components:
  schemas:
    People:
      required:
        - surname
        - name
        - address
      type: object
      properties:
        surname:
          type: string
          example: Иванов
        name:
          type: string
          example: Иван
        patronymic:
          type: string
          example: Иванович
        address:
          type: string
          example: г. Москва, ул. Ленина, д. 5, кв. 1

time_tracker/.env:

DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=db
DB_HOST=10.0.2.2
DB_PORT=5432

time_tracker/migrations/001_initial_setup.up.sql:

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    passport_number VARCHAR(10) NOT NULL,
    name VARCHAR(100),
    surname VARCHAR(100),
    patronymic VARCHAR(100),
    address VARCHAR(255)
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP
);

time_tracker/migrations/001_initial_setup.down.sql:

DROP TABLE tasks;
DROP TABLE users;

time_tracker/api/handler.go:

package api

import (
	"encoding/json"
	"time"
	"unicode"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"time_tracker/model"
	"time_tracker/database"
	log "github.com/sirupsen/logrus"
	"time_tracker/out/go"
	"context"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

// GetUsers handles the GET /users endpoint.
// It returns a list of users.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: []User
//       400: badRequestError
//       500: internalServerError
func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Debug("GetUsers called")
	db, err := database.NewDB()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error connecting to the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	filter := r.URL.Query().Get("filter")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	}

	users, err := database.GetUsers(db, filter, page)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error getting users from the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// GetUser handles the GET /users/{id} endpoint.
// It returns the user with the given ID.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: User
//       400: badRequestError
//       500: internalServerError
func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Debug("GetUser called")
	db, err := database.NewDB()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error connecting to the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	user, err := database.GetUser(db, params["id"])
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error getting user from the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// CreateUser handles the POST /users endpoint.
// It creates a new user.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: User
//       400: badRequestError
//       500: internalServerError
func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Debug("CreateUser called")
	db, err := database.NewDB()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error connecting to the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error decoding user data")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if !isValidPassportNumber(user.PassportNumber) {
		log.WithFields(log.Fields{
			"passportNumber": user.PassportNumber,
		}).Error("Invalid passport number")
		http.Error(w, "Invalid passport number", http.StatusBadRequest)
		return
	}

	cfg := openapi.NewConfiguration()
	c := openapi.NewAPIClient(cfg)

	passportSerie, err := strconv.ParseInt(user.PassportSerie, 10, 32)
	passportNumber, err := strconv.ParseInt(user.PassportNumber, 10, 32)

	info, _, err := c.DefaultAPI.InfoGet(context.Background()).PassportSerie(int32(passportSerie)).PassportNumber(int32(passportNumber)).Execute()
if err != nil {
	log.WithFields(log.Fields{
		"error": err,
	}).Warn("Error making request to Swagger API")
} else {
	user.Name = info.Name
	user.Surname = info.Surname
	if info.Patronymic != nil {
		user.Patronymic = *info.Patronymic
	}
	user.Address = info.Address
}

	err = database.CreateUser(db, &user)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error creating user in the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func isValidPassportNumber(passportNumber string) bool {
	if len(passportNumber) != 10 {
		return false
	}
	for _, char := range passportNumber {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

// DeleteUser handles the DELETE /users/{id} endpoint.
// It deletes the user with the given ID.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: successResponse
//       400: badRequestError
//       500: internalServerError
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Debug("DeleteUser called")
	db, err := database.NewDB()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error connecting to the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	err = database.DeleteUser(db, params["id"])
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error deleting user from the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

// UpdateUser handles the PUT /users/{id} endpoint.
// It updates the user with the given ID.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: User
//       400: badRequestError
//       500: internalServerError
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Debug("UpdateUser called")
	db, err := database.NewDB()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error connecting to the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	var user model.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error decoding user data")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = database.UpdateUser(db, params["id"], &user)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error updating user in the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// StartTask handles the POST /users/{id}/start endpoint.
// It starts a new task for the user with the given ID.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: successResponse
//       400: badRequestError
//       500: internalServerError
func StartTask(w http.ResponseWriter, r *http.Request) {
	log.Debug("StartTask called") 
	db, err := database.NewDB()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error connecting to the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	err = database.StartTask(db, params["id"])
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error starting task in the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

// EndTask handles the POST /users/{id}/end endpoint.
// It ends the current task for the user with the given ID.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: successResponse
//       400: badRequestError
//       500: internalServerError
func EndTask(w http.ResponseWriter, r *http.Request) {
	log.Debug("EndTask called")
	db, err := database.NewDB()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error connecting to the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	err = database.EndTask(db, params["id"])
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error ending task in the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

// GetTasksByUserAndPeriod handles the GET /users/{id}/tasks endpoint.
// It returns the tasks for the user with the given ID that were started and ended within the given period.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200: []Task
//       400: badRequestError
//       500: internalServerError
func GetTasksByUserAndPeriod(w http.ResponseWriter, r *http.Request) {
	log.Debug("GetTasksByUserAndPeriod called")
	db, err := database.NewDB()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error connecting to the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	start, err := time.Parse(time.RFC3339, r.URL.Query().Get("start"))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	end, err := time.Parse(time.RFC3339, r.URL.Query().Get("end"))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	tasks, err := database.GetTasksByUserAndPeriod(db, params["id"], start, end)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error getting tasks from the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

time_tracker/api/router.go:

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

time_tracker/api/server.go:

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

time_tracker/model/user.go:

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

time_tracker/config/config.go:

package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func GetConfig() Config {
	log.Debug("Loading .env file")
	err := godotenv.Load()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error loading .env file")
	}

	log.Debug("Successfully loaded .env file")

	return Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
	}
}

time_tracker/database/db.go:

package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"time_tracker/config"
	"time_tracker/model"
	"time"
	"unicode"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func Migrate() error {
	log.Debug("Starting database migration")
	db, err := NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	tableCheck := `SELECT EXISTS (
		SELECT FROM information_schema.tables 
		WHERE  table_schema = 'public'
		AND    table_name   = 'users'
	);`
	var exists bool
	err = db.QueryRow(tableCheck).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		log.Info("Table already exists, skipping migration")
		return nil
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	log.Info("Database migration completed")

	return nil
}

func NewDB() (*sql.DB, error) {
	log.Debug("Connecting to the database")
	cfg := config.GetConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Debug("Successfully connected to the database")

	return db, nil
}

func GetUsers(db *sql.DB, filter string, page int) ([]model.User, error) {
    log.Debug("Fetching users from the database")
    offset := (page - 1) * 10
	
    query := `SELECT * FROM users WHERE name LIKE $1 OR surname LIKE $1 OR patronymic LIKE $1 OR address LIKE $1 LIMIT 10 OFFSET $2`
    
    rows, err := db.Query(query, "%" + filter + "%", offset)
    if err != nil {
        log.WithFields(log.Fields{
            "error": err,
        }).Error("Error executing query")
        return nil, err
    }
    defer rows.Close()

    var users []model.User
    for rows.Next() {
        var user model.User
        err = rows.Scan(&user.ID, &user.PassportNumber, &user.Name, &user.Surname, &user.Patronymic, &user.Address)
        if err != nil {
            log.WithFields(log.Fields{
                "error": err,
            }).Error("Error scanning row")
            return nil, err
        }
        users = append(users, user)
    }

    if err = rows.Err(); err != nil {
        log.WithFields(log.Fields{
            "error": err,
        }).Error("Error fetching rows")
        return nil, err
    }

    log.Debug("Successfully fetched users from the database")

    return users, nil
}

func GetUser(db *sql.DB, id string) (model.User, error) {
	log.Debug("Fetching user from the database")
	var user model.User
	err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.PassportNumber, &user.Name, &user.Surname, &user.Patronymic, &user.Address)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id": id,
		}).Error("Error getting user from the database")
		return model.User{}, err
	}

	log.Debug("Successfully fetched user from the database")

	return user, nil
}

func CreateUser(db *sql.DB, user *model.User) error {
	log.Debug("Creating user in the database")
	if !isValidPassportNumber(user.PassportNumber) {
		log.WithFields(log.Fields{
			"passportNumber": user.PassportNumber,
		}).Error("Invalid passport number")
		return fmt.Errorf("invalid passport number: %s", user.PassportNumber)
	}

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE passport_number = $1)", user.PassportNumber).Scan(&exists)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error checking passport number in the database")
		return err
	}
	if exists {
		log.WithFields(log.Fields{
			"passportNumber": user.PassportNumber,
		}).Error("Passport number already exists in the database")
		return fmt.Errorf("passport number: %s already exists", user.PassportNumber)
	}

	_, err = db.Exec("INSERT INTO users (passport_number, name, surname, patronymic, address) VALUES ($1, $2, $3, $4, $5)", user.PassportNumber, user.Name, user.Surname, user.Patronymic, user.Address)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error creating user in the database")
		return err
	}

	log.Debug("Successfully created user in the database")

	return nil
}

func DeleteUser(db *sql.DB, id string) error {
	log.Debug("Deleting user from the database")
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id": id,
		}).Error("Error deleting user from the database")
		return err
	}

	log.Debug("Successfully deleted user from the database")

	return nil
}

func UpdateUser(db *sql.DB, id string, user *model.User) error {
	log.Debug("Updating user in the database")
	if !isValidPassportNumber(user.PassportNumber) {
		log.WithFields(log.Fields{
			"passportNumber": user.PassportNumber,
		}).Error("Invalid passport number")
		return fmt.Errorf("invalid passport number: %s", user.PassportNumber)
	}

	_, err := db.Exec("UPDATE users SET passport_number = $1, name = $2, surname = $3, patronymic = $4, address = $5 WHERE id = $6", user.PassportNumber, user.Name, user.Surname, user.Patronymic, user.Address, id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id": id,
			"user": user,
		}).Error("Error updating user in the database")
		return err
	}

	log.Debug("Successfully updated user in the database")

	return nil
}

func isValidPassportNumber(passportNumber string) bool {
	if len(passportNumber) != 10 {
		return false
	}
	for _, char := range passportNumber {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func StartTask(db *sql.DB, id string) error {
	log.Debug("Starting task in the database")
	_, err := db.Exec("INSERT INTO tasks (user_id, start_time) VALUES ($1, $2)", id, time.Now())
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id": id,
		}).Error("Error starting task in the database")
		return err
	}

	log.Debug("Successfully started task in the database")

	return nil
}

func EndTask(db *sql.DB, id string) error {
	log.Debug("Ending task in the database")
	_, err := db.Exec("UPDATE tasks SET end_time = $1 WHERE user_id = $2 AND end_time IS NULL", time.Now(), id)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"id": id,
		}).Error("Error ending task in the database")
		return err
	}

	log.Debug("Successfully ended task in the database")

	return nil
}

func GetTasksByUserAndPeriod(db *sql.DB, userID string, start, end time.Time) ([]model.Task, error) {
	log.Debug("Fetching tasks by user and period from the database")
	rows, err := db.Query("SELECT * FROM tasks WHERE user_id = $1 AND start_time >= $2 AND end_time <= $3", userID, start, end)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"userID": userID,
			"start": start,
			"end": end,
		}).Error("Error getting tasks from the database")
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err = rows.Scan(&task.ID, &task.UserID, &task.StartTime, &task.EndTime)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Error scanning task row")
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error fetching task rows")
		return nil, err
	}

	log.Debug("Successfully fetched tasks by user and period from the database")

	return tasks, nil
}

time_tracker/middleware/middleware.go:

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

time_tracker/utils/utils.go:

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

time_tracker/time_tracker.yaml:

openapi: 3.0.3
info:
  title: Time Tracker API
  version: 1.0.0
paths:
  /users:
    get:
      summary: Получение данных пользователей
      parameters:
        - name: filter
          in: query
          schema:
            type: string
        - name: page
          in: query
          schema:
            type: integer
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/User'
    post:
      summary: Добавление нового пользователя
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/User'
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/User'
  /users/{id}:
    get:
      summary: Получение данных пользователя по ID
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/User'
    put:
      summary: Изменение данных пользователя
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/User'
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                $ref: '#/components/schemas/User'
    delete:
      summary: Удаление пользователя
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                type: object
                properties:
                  result:
                    type: string
  /users/{id}/start:
    post:
      summary: Начать отсчет времени по задаче для пользователя
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                type: object
                properties:
                  result:
                    type: string
  /users/{id}/end:
    post:
      summary: Закончить отсчет времени по задаче для пользователя
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                type: object
                properties:
                  result:
                    type: string
  /users/{id}/tasks:
    get:
      summary: Получение трудозатрат по пользователю за период
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
        - name: start
          in: query
          schema:
            type: string
            format: date-time
        - name: end
          in: query
          schema:
            type: string
            format: date-time
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: '#/components/schemas/Task'
components:
  schemas:
    User:
      type: object
      properties:
        id:
          type: string
        passportNumber:
          type: string
        passportSerie:
          type: string
        name:
          type: string
        surname:
          type: string
        patronymic:
          type: string
        address:
          type: string
    Task:
      type: object
      properties:
        id:
          type: string
        userId:
          type: string
        startTime:
          type: string
          format: date-time
        endTime:
          type: string
          format: date-time

time_tracker/README.md:

# Time Tracker API

## Описание

Time Tracker API - это RESTful API, разработанный на Go. Он предоставляет функциональность для отслеживания времени, потраченного пользователями на различные задачи.

## Особенности

- Получение списка всех пользователей
- Получение информации о конкретном пользователе
- Создание нового пользователя
- Удаление пользователя
- Обновление информации о пользователе
- Начало отсчета времени по задаче для пользователя
- Завершение отсчета времени по задаче для пользователя
- Получение списка задач пользователя за определенный период
- Получение общего времени, потраченного пользователем на задачи за определенный период

## Технологии

- Язык программирования: Go
- База данных: PostgreSQL
- Библиотеки: Gorilla Mux (для маршрутизации), pq (для взаимодействия с PostgreSQL)

## Конфигурация

Проект использует файл `.env` для хранения конфигурационных данных. Это позволяет легко изменять конфигурацию без необходимости изменять код.

Вот пример того, как может выглядеть ваш файл `.env`:

```env
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
DB_HOST=your_database_host
DB_PORT=your_database_port
```

Замените your_database_user, your_database_password, your_database_name, your_database_host и your_database_port на соответствующие значения для вашей базы данных.

## Установка и запуск

1. Установите Go и PostgreSQL, если у вас их еще нет.
2. Docker (для Swagger UI).
3. Клонируйте этот репозиторий.
4. Перейдите в каталог проекта.
5. Установите зависимости с помощью команды `go get`.
6. Создайте базу данных в PostgreSQL и обновите файл `.env` с правильными данными для подключения к базе данных.
7. Запустите проект с помощью команды `go run main.go`. Это запустит сервер на порту 8080.

## Использование Swagger UI для просмотра документации API
Запустите Swagger UI с помощью Docker:

1. Запустите Swagger UI для вашего API тайм-трекера:

```
docker run -p 80:8080 -e SWAGGER_JSON=/foo/time_tracker.yaml -v /путь/к/вашему/каталогу/time_tracker:/foo swaggerapi/swagger-ui
```

Запустите Swagger UI для внешнего API:
```
docker run -p 81:8080 -e SWAGGER_JSON=/foo/swagger.yaml -v /путь/к/вашему/каталогу/time_tracker:/foo swaggerapi/swagger-ui
```

В этих командах замените /путь/к/вашему/каталогу/time_tracker на путь к каталогу проекта на вашем компьютере.

2. Откройте веб-браузер и перейдите по адресу http://localhost:80 для просмотра документации API тайм-трекера и http://localhost:81 для просмотра документации внешнего API.(Введите /swagger/swagger.yaml в поле для URL и нажмите кнопку “Explore”.)

## Использование API

API поддерживает следующие HTTP-запросы:

- `GET /users`: Возвращает список всех пользователей.
- `GET /users/{id}`: Возвращает пользователя с указанным ID.
- `POST /users/{id}`: Создает нового пользователя с указанным ID. Тело запроса должно содержать данные пользователя в формате JSON.
- `DELETE /users/{id}`: Удаляет пользователя с указанным ID.
- `PUT /users/{id}`: Обновляет пользователя с указанным ID. Тело запроса должно содержать обновленные данные пользователя в формате JSON.


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
	"database/sql"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func getDB(w http.ResponseWriter) (*sql.DB, error) {
	db, err := database.NewDB()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error connecting to the database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return nil, err
	}
	return db, nil
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
	db, err := getDB(w)
	if err != nil {
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
	db, err := getDB(w)
	if err != nil {
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
	db, err := getDB(w)
	if err != nil {
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
	db, err := getDB(w)
	if err != nil {
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
	db, err := getDB(w)
	if err != nil {
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
	db, err := getDB(w)
	if err != nil {
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
	db, err := getDB(w)
	if err != nil {
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
	db, err := getDB(w)
	if err != nil {
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



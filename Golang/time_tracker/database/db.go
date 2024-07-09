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
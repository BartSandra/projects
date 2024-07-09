package test

import (
	"testing"
	"time"
	"time_tracker/database"
	"time_tracker/model"
)

func TestGetUsers(t *testing.T) {
	db, err := database.NewDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	users, err := database.GetUsers(db, "", 1)
	if err != nil {
		t.Fatalf("Failed to get users: %v", err)
	}

	expectedNumOfUsers := 10
	if len(users) != expectedNumOfUsers {
		t.Errorf("Expected %d users, but got %d", expectedNumOfUsers, len(users))
	}
}

func TestCreateUser(t *testing.T) {
	db, err := database.NewDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	user := &model.User{
		PassportNumber: "9497567890",
		Name:           "Testt_Name",
		Surname:        "Testt_Surname",
	}

	err = database.CreateUser(db, user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}
}

func TestStartTask(t *testing.T) {
	db, err := database.NewDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	err = database.StartTask(db, "2")
	if err != nil {
		t.Fatalf("Failed to start task: %v", err)
	}
}

func TestEndTask(t *testing.T) {
	db, err := database.NewDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
	err = database.EndTask(db, "2")
	if err != nil {
		t.Fatalf("Failed to end task: %v", err)
	}
}

func TestGetTasksByUserAndPeriod(t *testing.T) {
	db, err := database.NewDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	start := time.Now().Add(-24 * time.Hour)
	end := time.Now()

	var tasks []model.Task
	tasks, err = database.GetTasksByUserAndPeriod(db, "3", start, end)
	if err != nil {
		t.Fatalf("Failed to get tasks: %v", err)
	}

	expectedNumOfTasks := 0
	if len(tasks) != expectedNumOfTasks {
		t.Errorf("Expected %d tasks, but got %d", expectedNumOfTasks, len(tasks))
	}
}

func TestUpdateUser(t *testing.T) {
	db, err := database.NewDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	user := &model.User{
		PassportNumber: "9293567890",
		Name:           "Test_Name_Updated",
		Surname:        "Test_Surname_Updated",
	}

	err = database.UpdateUser(db, "11", user)
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	updatedUser, err := database.GetUser(db, "11")
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	if updatedUser.Name != "Test_Name_Updated" || updatedUser.Surname != "Test_Surname_Updated" {
		t.Errorf("Expected user name and surname to be updated, but got %v", updatedUser)
	}
}

func TestDeleteUser(t *testing.T) {
	db, err := database.NewDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	err = database.DeleteUser(db, "10")
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	_, err = database.GetUser(db, "10")
	if err == nil {
		t.Errorf("Expected error when getting deleted user, but got nil")
	}
}

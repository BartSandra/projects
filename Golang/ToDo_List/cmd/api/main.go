package main

import (
    "log"
    "net/http"
    "os"
    "todo-app/internal/todo"
    "todo-app/pkg/db"

    _ "todo-app/docs" // подключение документации swagger
    httpSwagger "github.com/swaggo/http-swagger"
)

// @title Todo List API
// @description API для управления списком задач.
// @version 1.0.0
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
    cfg, err := db.NewConfig()
    if err != nil {
        log.Fatalf("Error loading database configuration: %v", err)
    }

    database, err := db.Init(cfg)
    if err != nil {
        log.Fatalf("Error initializing database: %v", err)
    }
    defer database.Close()

    todoHandler := todo.NewHandler(todo.NewService(todo.NewRepository(database)))

    http.HandleFunc("/todos", todoHandler.HandleTodos)
    http.HandleFunc("/swagger/", httpSwagger.WrapHandler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server is starting on port %s...", port)
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatalf("Could not start server: %v", err)
    }
}

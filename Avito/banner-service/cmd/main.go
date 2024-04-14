package main

import (
	"banner-service/pkg/config"
	"banner-service/pkg/db"
	"banner-service/pkg/handlers"
	"banner-service/pkg/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Загрузка конфигурации из файла .env или окружения
	cfg := config.Load()

	// Инициализация подключения к базе данных
	dbConn, err := db.Initialize(cfg)
	if err != nil {
		log.Fatal("Could not initialize database: ", err)
	}
	defer dbConn.Close()

	// Создание роутера с помощью gorilla/mux
	router := mux.NewRouter()

	// Применение middleware для аутентификации JWT
	router.Use(middleware.JWTAuthentication(cfg))

	// Регистрация обработчиков для различных эндпоинтов
	router.HandleFunc("/signin", handlers.Signin).Methods("POST")
	router.HandleFunc("/user_banner", handlers.GetUserBanner(dbConn)).Methods("GET")
	router.HandleFunc("/banner", handlers.GetBanners(dbConn)).Methods("GET")
	router.HandleFunc("/banner", handlers.CreateBanner(dbConn)).Methods("POST")
	router.HandleFunc("/banner/{id}", handlers.UpdateBanner(dbConn)).Methods("PATCH")
	router.HandleFunc("/banner/{id}", handlers.DeleteBanner(dbConn)).Methods("DELETE")

	// Настройка и запуск HTTP сервера
	log.Printf("Starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

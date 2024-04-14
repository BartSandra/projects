package server

import (
    "banner-service/pkg/config"
   // "banner-service/pkg/db"
    "banner-service/pkg/handlers"
    "banner-service/pkg/middleware"
    "database/sql"
    "github.com/gorilla/mux"
)

func SetupRouter(cfg *config.Config, dbConn *sql.DB) *mux.Router {
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

    return router
}
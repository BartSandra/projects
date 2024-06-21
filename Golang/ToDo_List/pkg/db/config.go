package db

import (
    "errors"
    "os"
)

type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

func NewConfig() (*Config, error) {
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    sslMode := os.Getenv("DB_SSLMODE")

    if host == "" || port == "" || user == "" || password == "" || dbName == "" || sslMode == "" {
        return nil, errors.New("one or more required environment variables are missing")
    }

    return &Config{
        Host:     host,
        Port:     port,
        User:     user,
        Password: password,
        DBName:   dbName,
        SSLMode:  sslMode,
    }, nil
}

package logger

import (
	"log"
	"os"
)

var Logger *log.Logger

func Init() {
	Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Error(msg string) {
	Logger.SetPrefix("ERROR: ")
	Logger.Println(msg)
}

func Info(msg string) {
	Logger.SetPrefix("INFO: ")
	Logger.Println(msg)
}

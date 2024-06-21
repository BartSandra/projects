package main

import (
	//"log"
	"test/internal/app"
	"test/internal/config"
)

func main() {
	config.Load()
	app.Start()
}



package main

import (
	"web-booking/services/logx"
	"web-booking/src/routes"

	"github.com/joho/godotenv"
)

func main() {

	logx.Init()
	if err := godotenv.Load(); err != nil {
		logx.Warnf("Error loading .env file")
	}

	routes.Routes()
}

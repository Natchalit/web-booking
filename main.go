package main

import (
	"web-booking/services/logx"
	"web-booking/src/routes"
)

func main() {

	logx.Init()

	routes.Routes()
}

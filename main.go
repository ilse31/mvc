package main

import (
	"mvc/config"
	m "mvc/middleware"
	"mvc/routes"
)

func main() {
	config.InitDB("")
	echoApp := routes.Routes()
	m.LogMiddlewareInit(echoApp)
	echoApp.Start(":8000")

}

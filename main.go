package main

import (
	route "github.com/Wong801/gin-api/src/routes"
)

func main() {
	server := route.InitRoutes()

	server.Run(":5000")
}

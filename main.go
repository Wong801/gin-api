package main

import (
	"github.com/Wong801/gin-api/src/db"
	route "github.com/Wong801/gin-api/src/routes"
)

func main() {
	server := route.InitRoutes()
	DB := db.InitDB()

	DB.MigrateModels()
	server.Run(":5000")
}

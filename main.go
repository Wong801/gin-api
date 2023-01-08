package main

import (
	"github.com/Wong801/gin-api/src/config"
	"github.com/Wong801/gin-api/src/db"
	route "github.com/Wong801/gin-api/src/routes"
	seeder "github.com/Wong801/gin-api/src/seeders"
)

func main() {
	server := route.InitRoutes()
	DB := db.InitDB()
	seeders := seeder.InitSeeder(DB)

	DB.MigrateModels()
	seeders.All()
	server.Run(":" + config.GetEnv("PORT", "8000"))
}

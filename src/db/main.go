package db

import (
	"log"

	"github.com/Wong801/gin-api/src/config"
	model "github.com/Wong801/gin-api/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type handler struct {
	Database *gorm.DB
}

func InitDB() handler {
	var db handler
	var err error
	config := config.GetDB()

	dbURL := "postgres://" + config.User + ":" + config.Pass + "@" + config.Host + config.Port + "/" + config.DB

	db.Database, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func (h handler) Close() error {
	postgresDB, _ := h.Database.DB()
	return postgresDB.Close()
}

func (h handler) MigrateModels() bool {
	h.Database.AutoMigrate(&model.User{})

	h.Close()
	return true
}

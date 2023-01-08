package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Wong801/gin-api/src/config"
	model "github.com/Wong801/gin-api/src/models"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Adapter struct {
	Database *gorm.DB
	Postgres *sql.DB
}

func Open(h *Adapter) error {
	var err error
	config := config.GetDB()
	dbURL := fmt.Sprintf("postgres://%s:%s@%s%s/%s?sslmode=%s", config.User, config.Pass, config.Host, config.Port, config.DB, config.SSL)

	h.Database, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return err
	}
	return nil
}

func InitDB() *Adapter {
	var db Adapter

	err := Open(&db)
	if err != nil {
		log.Fatalln(err)
	}

	db.Postgres, err = db.Database.DB()
	if err != nil {
		log.Fatalln(err)
	}

	return &db
}

func (h Adapter) Close() error {
	return h.Postgres.Close()
}

func (h Adapter) MigrateModels() {
	h.Database.AutoMigrate(&model.User{})
	h.Database.AutoMigrate(&model.Company{})
}

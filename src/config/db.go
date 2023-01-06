package config

import (
	entity "github.com/Wong801/gin-api/src/entities"
)

func GetDB() *entity.DBConnection {
	var dbConfig entity.DBConnection
	dbConfig.User = GetEnv("POSTGRES_USER", "postgres")
	dbConfig.Pass = GetEnv("POSTGRES_PASS", "")
	dbConfig.DB = GetEnv("POSTGRES_DB", "postgres")
	dbConfig.Host = GetEnv("POSTGRES_HOST", "127.0.0.1")
	dbConfig.Port = GetEnv("POSTGRES_PORT", ":5432")

	return &dbConfig
}

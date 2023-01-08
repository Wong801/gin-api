package service

import (
	"net/http"

	"github.com/Wong801/gin-api/src/db"
)

type RootService struct {
}

func InitRootService() *RootService {
	return &RootService{}
}

func (rs RootService) GetStats(db *db.Adapter) (int, interface{}, error) {
	stats := db.Postgres.Stats()

	err := db.Postgres.Close()

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, stats, nil
}

func (rs RootService) Ping() (int, map[string]string) {
	return http.StatusOK, map[string]string{
		"message": "pong",
	}
}

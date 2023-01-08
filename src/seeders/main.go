package seeder

import "github.com/Wong801/gin-api/src/db"

type seeder struct {
	db *db.Adapter
}

func InitSeeder(db *db.Adapter) *seeder {
	return &seeder{
		db,
	}
}

func (s seeder) All() error {
	s.SeedCompany()

	return s.db.Postgres.Close()
}

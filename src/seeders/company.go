package seeder

import (
	"errors"

	model "github.com/Wong801/gin-api/src/models"
	"gorm.io/gorm"
)

func (s seeder) SeedCompany() {
	if s.db.Database.Migrator().HasTable(&model.Company{}) {
		if err := s.db.Database.First(&model.Company{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			companies := []model.Company{
				{
					Name:     "Tavia House",
					Logo:     "/public/assets/images/company_logo/tavia_house.png",
					Link:     "https://taviadigitalsolusi.com/",
					State:    "Indonesia",
					City:     "Jakarta",
					Province: "Jakarta",
				},
				{
					Name:     "Rational Scale",
					Logo:     "/public/assets/images/company_logo/rational_scale.png",
					Link:     "https://rationalscale.it",
					State:    "Italy",
					City:     "Varese",
					Province: "Lombardy",
				},
				{
					Name:     "Pandatech",
					Logo:     "/public/assets/images/company_logo/pandatech.png",
					Link:     "https://pandatech.io",
					State:    "Indonesia",
					City:     "Jakarta",
					Province: "Jakarta",
				},
			}
			s.db.Database.Create(&companies)
		}
	}
}

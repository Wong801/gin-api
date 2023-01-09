package service

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/Wong801/gin-api/src/db"
	model "github.com/Wong801/gin-api/src/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CompanyService struct {
	db *db.Adapter
}

func InitCompanyService() *CompanyService {
	return &CompanyService{
		db: db.InitDB(),
	}
}

func (cs CompanyService) Get(id int) (int, *model.Company, error) {
	c := &model.Company{
		Id: id,
	}
	db.Open(cs.db)

	if err := cs.db.Database.First(&c).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, nil, errors.New("company data not found")
		}
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, c, nil
}

func (cs CompanyService) Search(name string) (int, []model.Company, error) {
	var err error
	var companies []model.Company
	db.Open(cs.db)

	if name != "" {
		err = cs.db.Database.Find(&companies, "name = ?", name).Error
	} else {
		err = cs.db.Database.Find(&companies).Error
	}

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, nil, errors.New("company data not found")
		}
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, companies, nil
}

func (cs CompanyService) Create(c model.Company) (int, *model.Company, error) {
	db.Open(cs.db)

	if err := cs.db.Database.Create(&c).Error; err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, &c, nil
}

func (cs CompanyService) Update(id int, c model.Company) (int, *model.Company, error) {
	c.Id = id
	db.Open(cs.db)

	if err := cs.db.Database.Save(&c).Error; err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, &c, nil
}

func (cs CompanyService) Delete(id int) (int, *model.Company, error) {
	var c model.Company
	db.Open(cs.db)

	if err := cs.db.Database.Clauses(clause.Returning{}).Delete(&c, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, nil, errors.New("company data not found")
		}
		return http.StatusInternalServerError, nil, err
	}

	if err := os.Remove(strings.TrimPrefix(c.Logo, "/")); err != nil && !os.IsNotExist(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, &c, nil
}

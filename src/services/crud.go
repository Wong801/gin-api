package service

import (
	"errors"
	"net/http"
	"reflect"
	"strings"
	"unicode"

	"github.com/Wong801/gin-api/src/db"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CRUDModel struct {
	record map[string]interface{}
}

type CRUDService struct {
	name  string
	db    *db.Adapter
	model interface{}
}

func InitCRUDService(n string, m interface{}) *CRUDService {
	return &CRUDService{
		name:  n,
		db:    db.InitDB(),
		model: m,
	}
}

func setNestedRecord(dst map[string]interface{}, v reflect.Value) map[string]interface{} {
	types := v.Type()
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i)
		if val.Type().Kind() == reflect.Struct {
			dst[types.Field(i).Name] = setNestedRecord(dst, val)
		} else {
			dst[types.Field(i).Name] = val
		}
	}
	return dst
}

func setRecord(v any) map[string]interface{} {
	record := make(map[string]interface{})
	values := reflect.ValueOf(v)
	types := values.Type()
	for i := 0; i < values.NumField(); i++ {
		name := types.Field(i).Name
		if name != "Timestamp" && name != "ID" {
			val := values.Field(i)
			if val.Type().Kind() == reflect.Struct {
				record[types.Field(i).Name] = setNestedRecord(record, val)
			} else {
				record[types.Field(i).Name] = val.Interface()
			}
		}
	}
	return record
}

func unsetRecord(record map[string]interface{}) {
	for key, val := range record {
		r := []rune(key)
		if unicode.IsUpper(r[0]) {
			newKey := string(append([]rune{unicode.ToLower(r[0])}, r[1:]...))
			if _, ok := record[newKey]; !ok {
				record[newKey] = val
			}
			delete(record, key)
		}
	}
}

func (crds CRUDService) Get(id int) (int, map[string]interface{}, error) {
	var data CRUDModel
	data.record = setRecord(crds.model)
	db.Open(crds.db)

	if err := crds.db.Database.Model(&crds.model).First(&data.record, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, nil, errors.New(crds.name + " data not found")
		}
		return http.StatusInternalServerError, nil, err
	}

	unsetRecord(data.record)
	return http.StatusOK, data.record, nil
}

func (crds CRUDService) Search(m any) (int, any, error) {
	db.Open(crds.db)

	if err := crds.db.Database.Model(crds.model).Find(&m).Error; err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, m, nil
}

func (crds CRUDService) Create(m any) (int, map[string]interface{}, error) {
	var data CRUDModel
	data.record = setRecord(m)
	db.Open(crds.db)

	if err := crds.db.Database.Model(&m).Create(&data.record).Error; err != nil {
		var perr *pgconn.PgError
		if ok := errors.As(err, &perr); ok {
			columnName := strings.Split(perr.ConstraintName, "_")
			return http.StatusConflict, nil, errors.New(columnName[len(columnName)-1] + " is already used")
		}
		return http.StatusInternalServerError, nil, err
	}

	unsetRecord(data.record)
	return http.StatusCreated, data.record, nil
}

func (crds CRUDService) Update(id int, m any) (int, map[string]interface{}, error) {
	var data CRUDModel
	data.record = setRecord(m)
	db.Open(crds.db)

	if err := crds.db.Database.Model(&m).Where("id = ?", id).Save(&data.record).Error; err != nil {
		return http.StatusInternalServerError, nil, err
	}

	unsetRecord(data.record)
	return http.StatusOK, data.record, nil
}

func (crds CRUDService) Delete(id int) (int, map[string]interface{}, error) {
	var data CRUDModel
	data.record = setRecord(crds.model)
	db.Open(crds.db)

	if err := crds.db.Database.Model(&crds.model).Clauses(clause.Returning{}).Delete(&data.record, id).Error; err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if _, ok := data.record["ID"]; !ok {
		return http.StatusNotFound, nil, errors.New(crds.name + " data not found")
	}

	unsetRecord(data.record)
	return http.StatusOK, data.record, nil
}

package service

import (
	"encoding/json"
	"fmt"

	entity "github.com/Wong801/gin-api/src/entities"
	"github.com/go-playground/validator/v10"
)

func makeTypeError(typeErr *json.UnmarshalTypeError, errs []entity.ValidationError) []entity.ValidationError {
	var reason string
	correctType := typeErr.Type.String()
	if correctType == "int" {
		reason = "value must be number"
	} else {
		reason = "value must be " + correctType
	}
	errs = append(errs, entity.ValidationError{Field: typeErr.Field, Reason: reason})
	return errs
}

func makeValidationError(validationErrs validator.ValidationErrors, errs []entity.ValidationError) []entity.ValidationError {

	for _, f := range validationErrs {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs = append(errs, entity.ValidationError{Field: f.Field(), Reason: err})
	}
	return errs
}

func MakeRequestError(err error) []entity.ValidationError {
	errs := []entity.ValidationError{}
	switch err.(type) {
	case *json.UnmarshalTypeError:
		return makeTypeError(err.(*json.UnmarshalTypeError), errs)
	case validator.ValidationErrors:
		return makeValidationError(err.(validator.ValidationErrors), errs)
	default:
		errs = append(errs, entity.ValidationError{Reason: "Unknown Error"})
	}

	return errs
}

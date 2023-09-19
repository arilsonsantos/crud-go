package validation

import (
	"encoding/json"
	"errors"
	"github.com/arilsonsantos/crud-go.git/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	en2 "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_tranlation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en2.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		en_tranlation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type.")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Cause{}
		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Cause{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Some fields are invalid.", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields.")
	}

}

package validation

import (
	"encoding/json"
	"errors"
	errors2 "github.com/arilsonsantos/crud-go.git/src/errors"
	"github.com/gin-gonic/gin/binding"
	en2 "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranlation "github.com/go-playground/validator/v10/translations/en"
)

var (
	//Validate = validator.New()
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en2.New()
		un := ut.New(en, en)
		transl, _ = un.GetTranslator("en")
		err := entranlation.RegisterDefaultTranslations(val, transl)
		if err != nil {
			return
		}
	}
}

func ValidateUserError(validationErr error) *errors2.ErrorDto {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return errors2.BadRequestError("Invalid field type.")
	} else if errors.As(validationErr, &jsonValidationError) {
		var errorsCauses []errors2.Cause
		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := errors2.Cause{
				Field:   e.Field(),
				Message: e.Translate(transl),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return errors2.BadRequestCauseError("Some fields are invalid.", errorsCauses)
	} else {
		return errors2.BadRequestError("Error trying to convert fields.")
	}

}

package validation

import (
	"encoding/json"
	"errors"

	"github.com/augustosang/api-go/src/configuration/rest_errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_error error) *rest_errors.Rest_Err {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_error, &jsonErr) {
		return rest_errors.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_error, &jsonValidationError) {
		errorsCauses := []rest_errors.Causes{}

		for _, e := range validation_error.(validator.ValidationErrors) {
			cause := rest_errors.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_errors.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_errors.NewBadRequestError("Error trying to convert fields")
	}
}

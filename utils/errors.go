package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func CheckErr(err error, message string) {
	if err != nil {
		log.Fatal(message)
	}
}

type validationError struct {
	StatusCode int               `json:"statusCode"`
	Success    bool              `json:"success"`
	Message    string            `json:"message"`
	Keys       map[string]string `json:"keys,omitempty"`
}

func Validate(w http.ResponseWriter, v interface{}) error {
	var result error = nil

	validate := validator.New()
	en := en.New()
	uni := ut.New(en, en)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")
	registerTanslations(validate, trans)

	//validate struct and send error message to the user
	err := validate.Struct(v)

	if err != nil {
		results := make(map[string]string)
		if _, ok := err.(*validator.InvalidValidationError); ok {
			json.NewEncoder(w).Encode(&validationError{
				StatusCode: 500,
				Message:    "The server has encountered an error",
				Keys:       nil,
			})
		}

		for _, e := range err.(validator.ValidationErrors) {
			results[e.Field()] = e.Translate(trans)
		}

		json.NewEncoder(w).Encode(&validationError{
			StatusCode: 400,
			Message:    "Invalid validation",
			Keys:       results,
		})
		result = errors.New("could not decode json")
	}
	return result
}

func registerTanslations(validate *validator.Validate, trans ut.Translator) {
	//validation for required field
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field, but is missing", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	// validation for email field
	validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "{0} is not a valid email", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})
}

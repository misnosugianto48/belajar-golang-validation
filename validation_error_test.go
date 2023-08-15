package belajargolangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidationError(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=6"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "misnocom",
		Password: "Misno3",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

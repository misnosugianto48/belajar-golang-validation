package belajargolangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestCrossField(t *testing.T) {
	type LoginRequest struct {
		Username        string `validate:"required,email"`
		Password        string `validate:"required,min=6"`
		ConfirmPassword string `validate:"required,min=5,eqfield=Password"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username:        "misno@example.com",
		Password:        "Misno123",
		ConfirmPassword: "Misno123",
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

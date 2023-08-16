package belajargolangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

type RegisterRequest struct {
	Username string `validate:"required"`
	Name     string `validate:"required"`
	Phone    string `validate:"required,numeric "`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func MustValidRegister(level validator.StructLevel) {

	registerRequest := level.Current().Interface().(RegisterRequest)
	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Name {

	} else {
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}

}

func TestStructValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegister, RegisterRequest{})

	request := RegisterRequest{
		Username: "",
		Name:     "",
		Email:    "",
		Phone:    "",
		Password: "",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

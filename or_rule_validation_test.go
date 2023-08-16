package belajargolangvalidation

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

// use pipe to or rule (|)
func TestOrRule(t *testing.T) {
	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	request := Login{
		Username: "1234",
		Password: "123bnfdshvgd",
	}

	validate := validator.New()
	err := validate.Struct(request)

	if err != nil {
		fmt.Println(err.Error())
	}
}

// Validation Cross field
func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("field not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCrossFieldValidator(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	type User struct {
		Username string `validate:"required,field_equals_ignore_case=Email|field_equals_ignore_case=Name"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	request := User{
		Username: "mail",
		Email:    "mail@example.com",
		Phone:    "43564624",
		Name:     "example",
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

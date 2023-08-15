package belajargolangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	validate := validator.New()
	if validate == nil {
		t.Error("Validate is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()

	//default string, dont use this
	//user:=""

	user1 := "ini string"

	err := validate.Var(user1, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateTwoVariable(t *testing.T) {
	validate := validator.New()

	password := "rahasia"
	confirmPassword := "salah"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()

	user1 := "123"

	err := validate.Var(user1, "required,numeric")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()

	user1 := "77221"

	err := validate.Var(user1, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}

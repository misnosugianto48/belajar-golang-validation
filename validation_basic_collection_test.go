package belajargolangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

// dive is tag to test collection of array, slice, and map
func TestBasicCollection(t *testing.T) {
	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"dive,required,min=1"`
	}

	validate := validator.New()
	request := User{
		Id:   "UserId",
		Name: "UserName",
		Addresses: []Address{
			{
				City:    "City",
				Country: "Country",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Sleeping",
			"",
			"Coding",
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

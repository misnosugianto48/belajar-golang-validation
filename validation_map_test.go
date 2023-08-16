package belajargolangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestMap(t *testing.T) {
	type Address struct {
		City string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"dive,required,min=1"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
	}

	validate := validator.New()
	request := User{
		Id:   "UserId",
		Name: "UserName",
		Addresses: []Address{
			{
				City: "City",
			},
			{
				City: "",
			},
		},
		Hobbies: []string{
			"Sleeping",
			"",
			"Coding",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SD 2",
			},
			"": {
				Name: "",
			},
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func TestBasicMap(t *testing.T) {
	type Address struct {
		City string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"dive,required,min=1"`
		Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys,dive"`
		Wallet    map[string]int    `validate:"dive,keys,required,endkeys,required,gt=1000"`
	}

	validate := validator.New()
	request := User{
		Id:   "UserId",
		Name: "UserName",
		Addresses: []Address{
			{
				City: "City",
			},
			{
				City: "",
			},
		},
		Hobbies: []string{
			"Sleeping",
			"",
			"Coding",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SD 2",
			},
			"": {
				Name: "",
			},
		},
		Wallet: map[string]int{
			"BCA":     1000000,
			"MANDIRI": 0,
			"":        1000,
		},
	}

	err := validate.Struct(request)
	if err != nil {
		fmt.Println(err.Error())
	}

}

package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Name struct {
	gorm.Model

	Name string `json:"name" validate:"nonzero"`
	Age  int    `json:"age" validate:"min=1"`

	// RG  string `json:"rg" validate: "len=9"`
	// CPF string `json:"cpf" validate: "len=11, regexp=^[0-9]*$"`
}

func CheckNamesStruct(n *Name) error {
	if err := validator.Validate(n); err != nil {

		return err
	}

	return nil

}

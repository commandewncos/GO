package model

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

// Key is the username

type Users struct {
	gorm.Model
	HashedPassword string `json:"password" validate:"nonzero"`
	Name           string `json:"name" validate:"nonzero"`
	Age            int    `json:"age" validate:"min=18"`
	Username       string `json:"username" validate:"nonzero"`
	SessionToken   string `json:"sessionToken"`
	CSRFToken      string `json:"csrftoken"`
}

func ValidateUsersStruct(u *Users) error {
	if err := validator.Validate(u); err != nil {

		return err
	}

	return nil

}

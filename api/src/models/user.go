package models

import (
	"api/src/support"
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        string    `json:"id,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Username  string    `json:"username,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
}

func (u *User) Validate(action string) error {
	err := u.checkEmptyFields(action)

	if err != nil {
		return err
	}

	u.trimFields()

	err = u.hashPassword()

	if err != nil {
		return err
	}

	return nil
}

func (u *User) checkEmptyFields(action string) error {
	if u.FirstName == "" {
		return errors.New("first name is required")
	}

	if u.LastName == "" {
		return errors.New("last name is required")
	}

	if u.Username == "" {
		return errors.New("username is required")
	}

	if action == "add" && u.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (u *User) trimFields() {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Username = strings.TrimSpace(u.Username)
}

func (u *User) hashPassword() error {
	hashedPassword, err := support.Hash(u.Password)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

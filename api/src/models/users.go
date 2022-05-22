package models

import (
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

func (u *User) Validate() error {
	err := u.checkEmptyFields()

	if err != nil {
		return err
	}

	u.trimFields()
	return nil
}

func (u *User) checkEmptyFields() error {
	if u.FirstName == "" {
		return errors.New("first name is required")
	}

	if u.LastName == "" {
		return errors.New("last name is required")
	}

	if u.Username == "" {
		return errors.New("username is required")
	}

	if u.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (u *User) trimFields() {
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
	u.Username = strings.TrimSpace(u.Username)
}

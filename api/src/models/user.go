package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/johnnyseubert/devbook/src/security"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

// step = "create" or "update"
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Nick == "" {
		return errors.New("nick is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if checkmail.ValidateFormat(user.Email) != nil {
		return errors.New(checkmail.ErrBadFormat.Error())
	}

	if step == "create" && user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "create" {
		hashedPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	return nil
}

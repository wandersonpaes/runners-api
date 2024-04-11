package userRunner

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/wandersonpaes/runners-api/internal/pkg/security"
)

func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	if err := user.format(stage); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("the name is needed and cannot be empty")
	}

	if user.Nick == "" {
		return errors.New("the nick is needed and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("the email is needed and cannot be empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("email is invalid")
	}

	if stage == "register" && user.Password == "" {
		return errors.New("the password is needed and cannot be empty")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		passwordWithHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passwordWithHash)
	}

	return nil
}

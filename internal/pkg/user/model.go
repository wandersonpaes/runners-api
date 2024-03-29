package user

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateOn time.Time `json:"createOn,omitempty"`
}

func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	user.format()
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

	if stage == "register" && user.Password == "" {
		return errors.New("the password is needed and cannot be empty")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}

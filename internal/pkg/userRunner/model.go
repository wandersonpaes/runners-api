package userRunner

import (
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

type NewPassword struct {
	New     string `json:"new"`
	Current string `json:"current"`
}

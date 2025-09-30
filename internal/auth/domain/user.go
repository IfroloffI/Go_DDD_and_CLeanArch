package domain

import (
	"errors"
	"strings"
)

// User — доменная сущность
type User struct {
	ID       string
	Email    string
	Password string // хэш
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email is required")
	}
	if !strings.Contains(u.Email, "@") {
		return errors.New("invalid email format")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

package models

import "fmt"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *Login) Validate() error {
	if l.Email == "" {
		return fmt.Errorf("missing field: email")
	}
	if len(l.Password) < 8 {
		return fmt.Errorf("invalid field: password length must be greater or equal to 8")
	}
	return nil
}
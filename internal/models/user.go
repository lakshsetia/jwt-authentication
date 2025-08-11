package models

import "fmt"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("missing field: name")
	}
	if u.Email == "" {
		return fmt.Errorf("missing field: email")
	}
	if len(u.Password) < 8 {
		return fmt.Errorf("invalid field: password length must be greater or equal to 8")
	}
	return nil
}
package db

import "github.com/lakshsetia/jwt-authentication/internal/models"

type DB interface {
	CreateUser(*models.User) (string, error)
	GetUserByID(id string) (models.User, error)
	AuthenticateUser(*models.Login) (models.User, error)
	DeleteUserById(string) error
}
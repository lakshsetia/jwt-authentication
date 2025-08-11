package pg

import (
	"database/sql"
	"fmt"

	"github.com/lakshsetia/jwt-authentication/internal/config"
	"github.com/lakshsetia/jwt-authentication/internal/models"
	"github.com/lakshsetia/jwt-authentication/internal/utils/password"
	_ "github.com/lib/pq"
)

type PG struct {
	db *sql.DB
}

func NewPG(config *config.Config) (*PG, error) {
	pg := config.PG
	user, password, dbname, port, host := pg.User, pg.Password, pg.DBName, pg.Port, pg.Host
	connectionStr := fmt.Sprintf("user=%v password=%v dbname=%v port=%v host=%v sslmode=disable", user, password, dbname, port, host)
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	_, err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)
	if err != nil {
		return nil, fmt.Errorf("failed to create extension: %w", err)
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users(
	id UUID PRIMARY KEY NOT NULL,
	name VARCHAR(256) NOT NULL,
	email VARCHAR(256) NOT NULL,
	hash_password VARCHAR(256) NOT NULL
	)`)
	if err != nil {
		return nil, fmt.Errorf("failed to create table users: %w", err)
	}
	return &PG{
		db: db,
	}, nil
}
func (pg *PG) CreateUser(user *models.User) (string, error) {
	err := pg.db.QueryRow("INSERT INTO users (id, name, email, hash_password) VALUES (uuid_generate_v4(), $1, $2, $3) RETURNING id", user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}
	return user.ID, nil
}
func (pg *PG) GetUserByID(id string) (models.User, error) {
	var user models.User
	err := pg.db.QueryRow("SELECT id, name, email, hash_password FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to retreive user: %w", err)
	}
	return user, nil
}
func (pg *PG) AuthenticateUser(login *models.Login) (models.User, error) {
	var user models.User
	err := pg.db.QueryRow("SELECT id, name, email, hash_password FROM users WHERE email=$1", login.Email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.User{}, fmt.Errorf("user not found with email=%v", login.Email)
		}
		return models.User{}, fmt.Errorf("failed to retrieve user: %w", err)
	}
	if !password.ComparePassword(user.Password, login.Password) {
		return models.User{}, fmt.Errorf("invalid password")
	}
	return user, nil
}
func (pg *PG) DeleteUserById(id string) error {
	_, err := pg.db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}
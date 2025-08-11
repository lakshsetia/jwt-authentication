package handlers

import (
	"net/http"

	"github.com/lakshsetia/jwt-authentication/internal/db"
	"github.com/lakshsetia/jwt-authentication/internal/jwt"
	"github.com/lakshsetia/jwt-authentication/internal/models"
	"github.com/lakshsetia/jwt-authentication/internal/utils/json"
	"github.com/lakshsetia/jwt-authentication/internal/utils/password"
	"github.com/lakshsetia/jwt-authentication/internal/utils/response"
)

func RegistrationHandler(db db.DB, key string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, response.MessageMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		var user models.User
		if err := json.ReadJSON(r, &user); err != nil {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, err.Error()), http.StatusBadRequest)
			return
		}
		if err := user.Validate(); err != nil {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, err.Error()), http.StatusBadRequest)
			return
		}
		hashPassword, err := password.GenerateHashPassword(user.Password)
		if err != nil {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, err.Error()), http.StatusInternalServerError)
			return 
		}
		user.Password = hashPassword
		user.ID, err = db.CreateUser(&user)
		if err != nil {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelDatabase, err.Error()), http.StatusInternalServerError)
			return
		}
		token, err := jwt.CreateToken(user.ID, key)
		if err != nil {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, response.MessageUnauthorized), http.StatusInternalServerError)
			return
		}
		json.WriteJSON(w, response.NewUserResponse(user.Name, user.Email, response.MessageRegistration, token), http.StatusCreated)	
	})
}

func LoginHandler(db db.DB, key string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, response.MessageMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		var login models.Login
		if err := json.ReadJSON(r, &login); err != nil {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, err.Error()), http.StatusBadRequest)
			return
		}
		if err := login.Validate(); err != nil {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, err.Error()), http.StatusBadRequest)
			return
		}
		user, err := db.AuthenticateUser(&login)
		if err != nil {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelDatabase, response.MessageUnauthorized), http.StatusUnauthorized)
			return
		}	
		token, err := jwt.CreateToken(user.ID, key)
		if err != nil {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, response.MessageUnauthorized), http.StatusInternalServerError)
			return
		}
		json.WriteJSON(w, response.NewUserResponse(user.Name, user.Email, response.MessageLogin, token), http.StatusOK)	
	})
}
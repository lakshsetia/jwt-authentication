package handlers

import (
	"net/http"
	"strings"

	"github.com/lakshsetia/jwt-authentication/internal/db"
	"github.com/lakshsetia/jwt-authentication/internal/jwt"
	"github.com/lakshsetia/jwt-authentication/internal/utils/json"
	"github.com/lakshsetia/jwt-authentication/internal/utils/response"
)

func AuthenticationHandler(db db.DB, key string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, response.MessageMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		tokenStr, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
		if !ok {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, response.MessageUnauthorized), http.StatusUnauthorized)
			return
		}
		claim, err := jwt.ValidateToken(tokenStr, key)
		if err != nil {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, response.MessageUnauthorized), http.StatusUnauthorized)
			return
		}
		user, err := db.GetUserByID(claim.ID)
		if err != nil {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelDatabase, err.Error()), http.StatusInternalServerError)
			return
		}
		json.WriteJSON(w, response.NewUserResponse(user.Name, user.Email, response.MessageAuthentication, tokenStr), http.StatusOK)
	})
}

func LogoutHandler(db db.DB, key string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, response.MessageMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
		tokenStr, ok := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
		if !ok {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, response.MessageUnauthorized), http.StatusUnauthorized)
			return
		}
		claim, err := jwt.ValidateToken(tokenStr, key)
		if err != nil {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelBackend, response.MessageUnauthorized), http.StatusUnauthorized)
			return
		}
		user, err := db.GetUserByID(claim.ID)
		if err != nil {
			json.WriteJSON(w, response.NewErrorResponse(response.LevelDatabase, err.Error()), http.StatusInternalServerError)
			return
		}
		newTokenStr := ""
		json.WriteJSON(w, response.NewUserResponse(user.Name, user.Email, response.MessageLogout, newTokenStr), http.StatusOK)
	})
}
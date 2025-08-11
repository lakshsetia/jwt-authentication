package server

import (
	"net/http"

	"github.com/lakshsetia/jwt-authentication/internal/handlers"
)

func (app *App) routes() http.Handler {
	handler := http.NewServeMux()
	handler.Handle("/user/register", handlers.RegistrationHandler(app.db, app.Key.HMACKey))
	handler.Handle("/user/auth", handlers.AuthenticationHandler(app.db, app.Key.HMACKey))
	handler.Handle("/user/login", handlers.LoginHandler(app.db, app.Key.HMACKey))
	handler.Handle("/user/logout", handlers.LogoutHandler(app.db, app.Key.HMACKey))
	return handler
}
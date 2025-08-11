package server

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lakshsetia/jwt-authentication/internal/config"
	"github.com/lakshsetia/jwt-authentication/internal/db"
	"github.com/lakshsetia/jwt-authentication/internal/db/pg"
)

type App struct {
	address string
	db 		db.DB
	config.Key
}

func NewApp(config *config.Config) (*App, error) {
	// connect database
	pg, err := pg.NewPG(config)
	if err != nil {
		return nil, err
	}
	return &App{
		address: config.Address,
		db: pg,
		Key: config.Key,
	}, nil
}
func (app *App) Run() {
	server := http.Server{
		Addr: app.address,
		Handler: app.routes(),
	}
	slog.Info("server starting at", slog.String("address", app.address))
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("failed to start server: %w", err)
		}
	}()
	<-done
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("failed to shutdown server: %w", err)
	}
}
package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}),
	}

	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("falha ao conectar ao Redis: %w", err)
	}

	fmt.Println("Iniciando servidor")

	err = server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("falha ao iniciar servidor: %w", err)
	}

	return nil

}

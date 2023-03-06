package app

import (
	"github.com/cucumberjaye/softline/configs"
	"github.com/cucumberjaye/softline/internal/app/handler"
	"github.com/cucumberjaye/softline/internal/app/repository/postgresdb"
	"github.com/cucumberjaye/softline/internal/app/service"
	"github.com/cucumberjaye/softline/pkg/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
)

type App struct {
	router *chi.Mux
}

func New() (*App, error) {
	err := configs.InitConfigs()
	if err != nil {
		return nil, err
	}

	db, err := postgres.New()
	if err != nil {
		return nil, err
	}

	repositories, err := postgresdb.New(db)
	if err != nil {
		return nil, err
	}

	services := service.New(repositories)
	handlers := handler.New(services)

	router := chi.NewRouter()

	router.Mount("/", handlers.InitRoutes())

	return &App{router: router}, nil
}

func (a *App) Run() error {
	log.Info().Msg("server running")

	return http.ListenAndServe(":8000", a.router)
}

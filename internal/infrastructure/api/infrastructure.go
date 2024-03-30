package api

import (
	"log"
	"pensatta/internal/infrastructure/postgres"

	"github.com/caarlos0/env/v9"
)

func (a *App) setupInfrastructure() {
	config, err := LoadEnvConfig[postgres.Config]()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db := postgres.NewGormPostgresClient(*config)
	userRepository := postgres.NewUserRepository(db)

	a.infrastructure = infrastructure{
		userRepository: userRepository,
	}
}

func LoadEnvConfig[T any]() (*T, error) {
	var config T
	return &config, env.Parse(&config)
}

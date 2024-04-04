package api

import (
	"pensatta/internal/core/exercise"
	"pensatta/internal/core/institution"
	"pensatta/internal/core/metrics"
	"pensatta/internal/core/user"
	"pensatta/internal/infrastructure/api/handler"
)

func (a *App) setupDependencies() {
	userService := user.NewService(a.infrastructure.userRepository)
	institutionService := institution.NewService(a.infrastructure.institutionRepository)
	exerciseService := exercise.NewService(a.infrastructure.exerciseRepository)
	metricsService := metrics.NewService(a.infrastructure.metricsRepository)

	a.depenedencies = depenedencies{
		registerHandler:    handler.NewRegisterHandler(userService),
		loginHandler:       handler.NewLoginHandler(userService),
		institutionHandler: handler.NewInstitutionHandler(institutionService),
		exerciseHandler:    handler.NewExerciseHandler(exerciseService),
		profileHandler:     handler.NewProfileHandler(exerciseService, metricsService),
	}
}

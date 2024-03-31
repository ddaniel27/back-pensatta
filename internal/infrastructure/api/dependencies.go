package api

import (
	"pensatta/internal/core/institution"
	"pensatta/internal/core/user"
	"pensatta/internal/infrastructure/api/handler"
)

func (a *App) setupDependencies() {
	userService := user.NewService(a.infrastructure.userRepository)
	institutionService := institution.NewService(a.infrastructure.institutionRepository)

	a.depenedencies = depenedencies{
		registerHandler:    handler.NewRegisterHandler(userService),
		loginHandler:       handler.NewLoginHandler(userService),
		institutionHandler: handler.NewInstitutionHandler(institutionService),
	}
}

package api

import (
	"pensatta/internal/core/user"
	"pensatta/internal/infrastructure/api/handler"
)

func (a *App) setupDependencies() {
	userService := user.NewService(a.infrastructure.userRepository)

	a.depenedencies = depenedencies{
		registerHandler: handler.NewRegisterHandler(userService),
	}
}

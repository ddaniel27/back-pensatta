package api

import (
	"context"
	"os"
	"os/signal"
	"pensatta/internal/core/ports/repositories"
	"pensatta/internal/infrastructure/api/handler"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type depenedencies struct {
	registerHandler    *handler.RegisterHandler
	loginHandler       *handler.LoginHandler
	institutionHandler *handler.InstitutionHandler
	exerciseHandler    *handler.ExerciseHandler
}

type infrastructure struct {
	userRepository        repositories.UserRepository
	institutionRepository repositories.InstitutionRepository
	exerciseRepository    repositories.ExerciseRepository
}

type App struct {
	Server *gin.Engine
	DB     *gorm.DB
	depenedencies
	infrastructure
}

func NewApp() *App {
	a := &App{}
	a.setupInfrastructure()
	a.setupDependencies()
	a.setupServer()

	return a
}

func (a *App) stopApp() {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
}

func (a *App) StartApp() {
	a.startServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	a.stopApp()
}

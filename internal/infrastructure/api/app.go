package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pensatta/internal/core/user"
	"pensatta/internal/infrastructure/api/handler"
	"time"

	"github.com/gin-gonic/gin"
)

type depenedencies struct {
	registerHandler *handler.RegisterHandler
}

type App struct {
	Server *gin.Engine
	depenedencies
}

func NewApp() *App {
	a := &App{}
	a.setupDependencies()
	a.setupServer()

	return a
}

func (a *App) setupDependencies() {
	userService := user.NewService()

	a.depenedencies = depenedencies{
		registerHandler: handler.NewRegisterHandler(userService),
	}
}

func (a *App) setupServer() {
	a.Server = gin.New()
	a.Server.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	baseGroup := a.Server.Group("/api")
	setupHealthCheckRoute(baseGroup)

	a.setupRoutes(baseGroup)
}

func (a *App) startServer() {
	if err := a.Server.Run(fmt.Sprintf(":%s", getPortFallback("PORT", "3000"))); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
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

func setupHealthCheckRoute(g *gin.RouterGroup) {
	g.GET("/health-check", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
}

func getPortFallback(env string, fallback string) string {
	port := os.Getenv(env)
	if port == "" {
		return fallback
	}
	return port
}

package api

import (
	"fmt"
	"log"
	"net/http"
	"pensatta/internal/infrastructure/api/validators"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func (a *App) setupServer() {
	a.Server = gin.New()
	a.Server.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("role", validators.ValidateRole)
	}

	baseGroup := a.Server.Group("/api")
	setupHealthCheckRoute(baseGroup)

	a.setupRoutes(baseGroup)
}

func (a *App) startServer() {
	if err := a.Server.Run(fmt.Sprintf(":%s", getPortFallback("PORT", "3000"))); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupHealthCheckRoute(g *gin.RouterGroup) {
	g.GET("/health-check", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
}

package api

import (
	"pensatta/internal/infrastructure/api/middlewares"

	"github.com/gin-gonic/gin"
)

func (a *App) setupRoutes(g *gin.RouterGroup) {
	registerGroup := g.Group("/register")
	registerGroup.Use(middlewares.AdminUser())
	registerGroup.POST("", a.registerHandler.CreateUser)
}

package api

import (
	"pensatta/internal/infrastructure/api/middlewares"

	"github.com/gin-gonic/gin"
)

func (a *App) setupRoutes(g *gin.RouterGroup) {
	registerGroup := g.Group("/register")
	registerGroup.Use(middlewares.AdminUser())
	registerGroup.POST("", a.registerHandler.CreateUser)

	loginGroup := g.Group("/login")
	loginGroup.POST("", middlewares.SetSession(), a.loginHandler.CreateLogin)
	loginGroup.GET("", middlewares.GetSession(), a.loginHandler.GetLogin)
}

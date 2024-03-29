package api

import "github.com/gin-gonic/gin"

func (a *App) setupRoutes(g *gin.RouterGroup) {
	registerGroup := g.Group("/register")
	registerGroup.POST("", a.registerHandler.CreateUser)
}

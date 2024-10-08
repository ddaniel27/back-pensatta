package api

import (
	"pensatta/internal/infrastructure/api/middlewares"

	"github.com/gin-gonic/gin"
)

func (a *App) setupRoutes(g *gin.RouterGroup) {
	registerGroup := g.Group("/register")
	// registerGroup.Use(middlewares.AdminUser())
	registerGroup.POST("", a.registerHandler.CreateUser)

	loginGroup := g.Group("/login")
	loginGroup.POST("", middlewares.SetSession(), a.loginHandler.CreateLogin)
	loginGroup.GET("", middlewares.GetSession(), a.loginHandler.GetLogin)

	logoutGroup := g.Group("/logout")
	logoutGroup.DELETE("", middlewares.DeleteSession())

	institutionGroup := g.Group("/institution")
	institutionGroup.Use(middlewares.AdminPermissions())
	institutionGroup.POST("", a.institutionHandler.CreateInstitution)
	institutionGroup.GET("", a.institutionHandler.GetInstitutions)
	institutionGroup.DELETE("/:id", a.institutionHandler.DeleteInstitution)

	exerciseGroup := g.Group("/exercise")
	exerciseGroup.Use(middlewares.GetSession())
	exerciseGroup.POST("", a.exerciseHandler.CreateExercise)

	profileGroup := g.Group("/profile")
	profileGroup.Use(middlewares.GetSession())
	profileGroup.GET("/exercises", a.profileHandler.GetAllExercisesForUser)
	profileGroup.GET("/metrics", a.profileHandler.GetMetricsForUser)
	profileGroup.GET("/resumen", a.profileHandler.GetResumenForUser)
}

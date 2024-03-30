package api

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"pensatta/internal/core/domain"
	"pensatta/internal/infrastructure/api/validators"

	"github.com/gin-contrib/sessions"
	gormsession "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type serverConfig struct {
	Port         string `env:"PORT"                   envDefault:"3000"`
	CookieSecret string `env:"COOKIE_SECRET,required"`
}

func (a *App) setupServer() {
	config, err := LoadEnvConfig[serverConfig]()
	if err != nil {
		log.Fatalf("Failed to load server config: %v", err)
	}

	sessionStore := gormsession.NewStore(a.DB, true, []byte(config.CookieSecret))
	gob.Register(domain.User{})

	a.Server = gin.New()
	a.Server.Use(
		gin.Recovery(),
		gin.Logger(),
		sessions.Sessions("pensatta-session", sessionStore),
	)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("role", validators.ValidateRole)
	}

	baseGroup := a.Server.Group("/api")
	setupHealthCheckRoute(baseGroup)

	a.setupRoutes(baseGroup)
}

func (a *App) startServer() {
	config, err := LoadEnvConfig[serverConfig]()
	if err != nil {
		log.Fatalf("Failed to load server config: %v", err)
	}

	if err := a.Server.Run(fmt.Sprintf(":%s", config.Port)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupHealthCheckRoute(g *gin.RouterGroup) {
	g.GET("/health-check", func(ctx *gin.Context) {
		ctx.Status(http.StatusOK)
	})
}

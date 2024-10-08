package handler

import (
	"pensatta/internal/core/domain"
	"pensatta/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	es services.ExerciseService
	ms services.MetricsService
	us services.UserService
}

func NewProfileHandler(
	es services.ExerciseService,
	ms services.MetricsService,
	us services.UserService,
) *ProfileHandler {
	return &ProfileHandler{
		es: es,
		ms: ms,
		us: us,
	}
}

func (ph *ProfileHandler) GetAllExercisesForUser(c *gin.Context) {
	ctx := c.Request.Context()
	user, ok := c.Get("user")
	if !ok {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	exercises, err := ph.es.GetExercisesByUserID(ctx, user.(domain.User).ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": "History retrieved", "history": exercises})
}

func (ph *ProfileHandler) GetMetricsForUser(c *gin.Context) {
	ctx := c.Request.Context()
	user, ok := c.Get("user")
	if !ok {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	metrics, err := ph.ms.GetMetricsByUserID(ctx, user.(domain.User).ID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		200,
		gin.H{
			"msg":                 "Metrics retrieved",
			"spider_values":       metrics.SpiderValues,
			"appropiation_values": metrics.AppropiationValues,
		},
	)
}

func (ph *ProfileHandler) GetResumenForUser(c *gin.Context) {
	ctx := c.Request.Context()
	user, ok := c.Get("user")
	if !ok {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	resumen, err := ph.us.GetUserProfileResumen(ctx, user.(domain.User))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(
		200,
		gin.H{
			"msg":              "Resumen retrieved",
			"institution_name": resumen["institution_name"],
			"last_login":       resumen["last_login"],
			"resumen":          resumen["resumen"],
		},
	)
}

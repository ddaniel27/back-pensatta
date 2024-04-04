package handler

import (
	"pensatta/internal/core/domain"
	"pensatta/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	es services.ExerciseService
}

func NewProfileHandler(es services.ExerciseService) *ProfileHandler {
	return &ProfileHandler{
		es: es,
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

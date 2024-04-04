package handler

import (
	"pensatta/internal/core/domain"
	"pensatta/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type ExerciseHandler struct {
	es services.ExerciseService
}

type exerciseBody struct {
	Score    float64 `json:"score"    binding:"required"`
	Time     uint64  `json:"time"     binding:"required"`
	Exercise uint64  `json:"exercise" binding:"required"`
	ID       uint64  `json:"id"`
}

func NewExerciseHandler(es services.ExerciseService) *ExerciseHandler {
	return &ExerciseHandler{
		es: es,
	}
}

func (eh *ExerciseHandler) CreateExercise(c *gin.Context) {
	var body exerciseBody
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, ok := c.Get("user")
	if !ok {
		c.JSON(400, gin.H{"error": "user not found"})
		return
	}

	if err := eh.es.
		CreateExercise(ctx, toDomainExercise(body), user.(domain.User).ID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "exercise created"})
}

func toDomainExercise(body exerciseBody) domain.Exercise {
	return domain.Exercise{
		ID:    body.Exercise,
		Score: body.Score,
		Time:  body.Time,
	}
}

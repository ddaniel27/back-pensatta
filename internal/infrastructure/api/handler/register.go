package handler

import (
	"pensatta/internal/core/domain"
	"pensatta/internal/core/ports/services"
	"strings"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	us services.UserService
}

func NewRegisterHandler(us services.UserService) *RegisterHandler {
	return &RegisterHandler{us: us}
}

func (uh *RegisterHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var user domain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	user.Role = strings.ToUpper(user.Role)

	if !validateRole(user.Role) {
		c.JSON(400, gin.H{"error": "Invalid role"})

		return
	}

	if err := uh.us.CreateUser(ctx, user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})

		return
	}

	c.JSON(201, gin.H{"message": "User created successfully"})
}

func validateRole(role string) bool {
	roles := []string{
		"STUDENT",
		"TEACHER",
		"COORDINATOR",
	}

	for _, r := range roles {
		if r == role {
			return true
		}
	}

	return false
}

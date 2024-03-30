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

type registerBody struct {
	FirstName       string `json:"first_name"       binding:"required"`
	LastName        string `json:"last_name"        binding:"required"`
	ListNumber      uint64 `json:"list_number"      binding:"required"`
	Role            string `json:"role"             binding:"required,role"`
	InstitutionCode string `json:"institution_code" binding:"required"`
	Password        string `json:"password"         binding:"required"`
}

func NewRegisterHandler(us services.UserService) *RegisterHandler {
	return &RegisterHandler{us: us}
}

func (uh *RegisterHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var user registerBody

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	userDomain := toUserDomain(user)
	username, err := uh.us.CreateUser(ctx, userDomain)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})

		return
	}

	c.JSON(201, gin.H{"username": username})
}

func toUserDomain(user registerBody) domain.User {
	return domain.User{
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		ListNumber:      user.ListNumber,
		Role:            strings.ToUpper(user.Role),
		InstitutionCode: user.InstitutionCode,
		Password:        user.Password,
	}
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

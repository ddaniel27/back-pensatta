package handler

import (
	"pensatta/internal/core/ports/services"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	us services.UserService
}

type loginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewLoginHandler(us services.UserService) *LoginHandler {
	return &LoginHandler{us: us}
}

func (lh *LoginHandler) CreateLogin(c *gin.Context) {
	var login loginBody
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})

		return
	}

	user, err := lh.us.ValidateCredentials(ctx, login.Username, login.Password)
	if err != nil {
		c.JSON(401, gin.H{"logged": false, "user": nil})

		return
	}

	c.Set("user", user)
}

func (lh *LoginHandler) GetLogin(c *gin.Context) {
	if user, ok := c.Get("user"); ok {
		c.JSON(200, gin.H{"logged": true, "user": user})

		return
	}
}

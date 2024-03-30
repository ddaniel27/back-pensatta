package middlewares

import (
	"pensatta/internal/core/domain"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AdminUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("pensatta")
		if user == nil || user.(domain.User).Role != "ADMIN" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()

			return
		}

		c.Next()
	}
}

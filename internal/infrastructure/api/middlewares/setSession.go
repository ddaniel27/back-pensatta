package middlewares

import (
	"fmt"
	"pensatta/internal/core/domain"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		c.Next()

		if v, ok := c.Get("user"); ok {
			user := v.(domain.User)
			user.Password = ""

			session.Set("pensatta-user", user)
			err := session.Save()
			if err != nil {
				c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to save session: %v", err)})
			}

			c.JSON(200, gin.H{"logged": true, "user": user})
		}
	}
}

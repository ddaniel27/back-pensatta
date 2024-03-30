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

			session.Set(domain.SessionUserKey, user)
			if err := session.Save(); err != nil {
				c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to save session: %v", err)})
				c.Abort()

				return
			}

			c.JSON(200, gin.H{"logged": true, "user": user})
		}
	}
}

func GetSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(domain.SessionUserKey)

		if user == nil {
			c.JSON(401, gin.H{"logged": false, "user": map[string]string{}})
			c.Abort()

			return
		}

		c.Set("user", user)

		c.Next()
	}
}

func DeleteSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Delete(domain.SessionUserKey)

		if err := session.Save(); err != nil {
			c.JSON(500, gin.H{"error": fmt.Sprintf("Failed to delete session: %v", err)})
			c.Abort()

			return
		}

		c.JSON(200, gin.H{"logged": false, "user": map[string]string{}})
	}
}

func AdminUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get(domain.SessionUserKey)

		if user == nil || user.(domain.User).Role != "ADMIN" {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()

			return
		}

		c.Next()
	}
}

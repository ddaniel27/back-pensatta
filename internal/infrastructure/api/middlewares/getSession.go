package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("pensatta-user")

		if user == nil {
			c.JSON(401, gin.H{"logged": false, "user": map[string]string{}})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}

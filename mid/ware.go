package mid

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"exchange/utils"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not token"})
			c.Abort()
			return
		}
		username, err := utils.ChakeToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not chked token"})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Next()
	}
}

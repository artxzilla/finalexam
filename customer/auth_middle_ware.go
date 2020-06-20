package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authMiddleWare(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}

package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func testHandler(c *gin.Context) {
	cus := &Customer{}

	c.JSON(http.StatusOK, cus)
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(authMiddleWare)

	r.GET("/test", testHandler)

	return r
}

func authMiddleWare(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}

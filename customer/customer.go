package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(authMiddleWare)

	r.POST("/customers", createCustomerHandler)
	r.GET("/customers/:id", getCustomerByIDHandler)
	r.GET("/customers", getAllCustomerHandler)
	r.PUT("/customers/:id", updateCustomerByIDHandler)
	r.DELETE("/customers/:id", deleteCustomerByIDHandler)

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

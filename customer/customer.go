package customer

import (
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

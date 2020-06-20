package customer

import (
	"fmt"
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
	fmt.Println("start middleware")

	c.Next()

	fmt.Println("end middleware")
}

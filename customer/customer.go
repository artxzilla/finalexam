package customer

import (
	"net/http"

	"github.com/artxzilla/finalexam/database"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(authMiddleWare)

	r.GET("/test", testHandler)
	r.POST("/customers", createCustomerHandler)
	r.GET("/customers/:id", getCustomerByIDHandler)
	r.GET("/customers", getAllCustomerHandler)

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

func testHandler(c *gin.Context) {
	cus := &Customer{}

	c.JSON(http.StatusOK, cus)
}

func createCustomerHandler(c *gin.Context) {
	cus := Customer{}
	if err := c.ShouldBindJSON(&cus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	row := database.GetInstance().QueryRow("INSERT INTO customers (name, email, status) values ($1, $2, $3) RETURNING id", cus.Name, cus.Email, cus.Status)

	err := row.Scan(&cus.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, cus)
}

func getCustomerByIDHandler(c *gin.Context) {
	id := c.Param("id")

	stmt, err := database.GetInstance().Prepare("SELECT id, name, email, status FROM customers where id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	res := stmt.QueryRow(id)

	cus := &Customer{}

	err = res.Scan(&cus.ID, &cus.Name, &cus.Email, &cus.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, cus)
}

func getAllCustomerHandler(c *gin.Context) {
	stmt, err := database.GetInstance().Prepare("SELECT id, name, email, status FROM customers")
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	results, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	customers := []*Customer{}
	for results.Next() {
		customer := &Customer{}

		err := results.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		customers = append(customers, customer)
	}

	c.JSON(http.StatusOK, customers)
}

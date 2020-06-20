package customer

import (
	"net/http"

	"github.com/artxzilla/finalexam/database"
	"github.com/gin-gonic/gin"
)

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

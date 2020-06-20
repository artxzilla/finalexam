package customer

import (
	"net/http"

	"github.com/artxzilla/finalexam/database"
	"github.com/gin-gonic/gin"
)

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

package customer

import (
	"net/http"

	"github.com/artxzilla/finalexam/database"
	"github.com/gin-gonic/gin"
)

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

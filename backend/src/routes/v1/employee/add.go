package employee

import (
	"context"
	"net/http"

	"github.com/dominus10/pos/db"
	"github.com/gin-gonic/gin"
)

func AddEmployee(ctx context.Context, q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {	
		var req struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		}

		_, err := q.AddEmployee(ctx,db.AddEmployeeParams{
			Name   					:"test1",
			RestaurantName 	:"admin",
			RoleName 				:"test",
			Email  					:req.Email,
			Crypt  					:req.Password,
		})

		if err != nil {
			c.Status(http.StatusInternalServerError)
		}

		c.Status(http.StatusCreated)
	}
}
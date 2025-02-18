package user

import (
	"context"
	"net/http"

	"github.com/dominus10/pos/db"
	"github.com/gin-gonic/gin"
)

func RegisterRestaurantOwner(ctx context.Context, q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Name    	string `json:"name" binding:"required"`
			Email	 		string `json:"email" binding:"required"`
			Password	string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
			_,e := q.RegisterRestaurantOwner(ctx,db.RegisterRestaurantOwnerParams{
			Column1: req.Name,
			Column2: req.Email,
			Crypt: req.Password,
		})
		if( e != nil){
			c.Status(http.StatusForbidden)
		}
		c.Status(http.StatusCreated)
	}
}
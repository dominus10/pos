package restaurant

import (
	"context"
	"net/http"

	"github.com/dominus10/pos/db"
	"github.com/gin-gonic/gin"
)

func InsertNewRestaurant(ctx context.Context, q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Name    string `json:"name" binding:"required"`
			Address string `json:"address" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
			_,e := q.InsertNewRestaurant(ctx,db.InsertNewRestaurantParams{
			Name: req.Name,
			Address: req.Address,
		})
		if( e != nil){
			panic("Cannot insert")
		}
		c.JSON(200, gin.H{
			"message":"Created",
		})
	}
}
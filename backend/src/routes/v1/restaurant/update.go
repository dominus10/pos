package restaurant

import (
	"context"
	"net/http"

	"github.com/dominus10/pos/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func UpdateRestaurant(ctx context.Context, q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			ID    	string `json:"id" binding:"required"`
			Name    string `json:"name" binding:"required"`
			Address string `json:"address" binding:"required"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		// Convert string to UUID
		parsedUUID, err := uuid.Parse(req.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
			return
		}

		// Convert to pgtype.UUID
		uuidValue := pgtype.UUID{
			Bytes: parsedUUID,
			Valid: true,
		}
		_,e := q.UpdateExistingRestaurant(ctx,db.UpdateExistingRestaurantParams{
			ID: uuidValue,
			Name: req.Name,
			Address: req.Address,
		})
		if( e != nil){
			panic("Cannot update!")
		}
		c.JSON(200, gin.H{
			"message":"Created",
		})
	}
}
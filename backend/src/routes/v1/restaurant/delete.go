package restaurant

import (
	"context"
	"net/http"

	"github.com/dominus10/pos/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func DeleteRestaurant(ctx context.Context, q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			ID    	string `json:"id" binding:"required"`
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

		i,e := q.DeleteRestaurant(ctx,uuidValue)

		if( e != nil){
			panic("Cannot delete!")
		}
		c.JSON(200, gin.H{
			"message":"Deleted",
			"data": gin.H{
				"name":i.Name,
				"address":i.Address,
			},
		})
	}
}
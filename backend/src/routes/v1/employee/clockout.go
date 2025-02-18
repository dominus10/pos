package employee

import (
	"context"
	"net/http"

	"github.com/dominus10/pos/db"
	"github.com/gin-gonic/gin"
)

// Clock-out API
func EmployeeClockOut(ctx context.Context, q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email    string `json:"email" binding:"required"`
		}
		// Execute clock-out query
		employee, err := q.EmployeeClockOut(ctx, req.Email)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot clock out or already clocked out"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":       "Clocked out successfully",
			"clock_out_time": employee.ClockOutTime,
		})
	}
}
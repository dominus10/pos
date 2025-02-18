package employee

import (
	"context"
	"net/http"
	"time"

	"github.com/dominus10/pos/db"
	"github.com/gin-gonic/gin"
)

// Clock-in API
func EmployeeClockIn(ctx context.Context, q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {	
		var req struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		// Bind JSON request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// Fetch employee details
		employee, err := q.GetEmployeeByEmail(ctx, req.Email)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid email or password"})
			return
		}

		// Check if already clocked in today
		today := time.Now().Truncate(24 * time.Hour)
		if employee.ClockInTime.Valid { // Ensure the timestamp exists
			clockInDate := employee.ClockInTime.Time.Truncate(24 * time.Hour) // Normalize for comparison
			if clockInDate.Equal(today) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Already clocked in today"})
				return
			}
		}

		// Execute clock-in query
		updatedEmployee, err := q.EmployeeClockIn(ctx,db.EmployeeClockInParams{
			Email: employee.Email,
			Crypt: req.Password,
		} )
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Clock-in failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":       "Clocked in successfully",
			"clock_in_time": updatedEmployee.ClockInTime,
		})
	}
}

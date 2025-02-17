package restaurant

import (
	"context"
	"net/http"

	"github.com/dominus10/pos/db"
	"github.com/dominus10/pos/src/security"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func UpdateRestaurant(ctx context.Context, q *db.Queries) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from JWT
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Get nonce from header
		nonce := c.GetHeader("X-Nonce")
		if nonce == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nonce missing"})
			return
		}

		// Validate Nonce (Prevent Replay Attacks)
		if !security.ValidateNonce(nonce) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Invalid nonce"})
			return
		}

		// Parse JSON request
		var req struct {
			ID      string `json:"id" binding:"required"`
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

		// Ensure the user owns the restaurant (Authorization)
		restaurant, err := q.GetRestaurant(ctx, uuidValue)
		if err != nil || restaurant.ID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
			return
		}

		// Update restaurant in DB
		_, err = q.UpdateExistingRestaurant(ctx, db.UpdateExistingRestaurantParams{
			ID:      uuidValue,
			Name:    req.Name,
			Address: req.Address,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot update restaurant"})
			return
		}

		// Return success response
		c.JSON(http.StatusOK, gin.H{
			"message": "Updated successfully",
		})
	}
}

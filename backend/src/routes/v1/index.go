package v1

import (
	"context"

	"github.com/dominus10/pos/db"
	"github.com/dominus10/pos/src/middleware"
	"github.com/dominus10/pos/src/routes/v1/restaurant"
	"github.com/dominus10/pos/src/routes/v1/user"
	"github.com/dominus10/pos/src/routes/v1/util"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine,ctx context.Context,q *db.Queries){
	v1 := r.Group("/v1")
	{
		v1.GET("/healthcheck",util.HealthCheckEndpoint)
		RestaurantRoutes(v1,ctx,q)
		UserRoutes(v1,ctx,q)
	}
}

func RestaurantRoutes(r *gin.RouterGroup,ctx context.Context,q *db.Queries){
	sub := r.Group("/restaurant", middleware.AuthMiddleware())
	{
		sub.POST("/register", restaurant.InsertNewRestaurant(ctx,q))
		sub.PUT("/register", restaurant.UpdateRestaurant(ctx,q))
		sub.DELETE("/register", restaurant.DeleteRestaurant(ctx,q))
	}
}

func UserRoutes(r *gin.RouterGroup,ctx context.Context,q *db.Queries){
	sub := r.Group("/user")
	{
		sub.POST("/register", user.RegisterRestaurantOwner(ctx,q))
	}
}

func UserSubRoutes(r *gin.RouterGroup,ctx context.Context,q *db.Queries){
	sub := r.Group("/guarded",middleware.AuthMiddleware())
	{
		sub.POST("/add", restaurant.InsertNewRestaurant(ctx,q))
	}
}
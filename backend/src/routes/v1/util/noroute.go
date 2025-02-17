package util

import "github.com/gin-gonic/gin"

func NoRouteFound(c *gin.Context){
	c.JSON(404, gin.H{
		"message":"Page Not Found",
	})
}
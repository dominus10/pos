package util

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HealthCheckEndpoint(c *gin.Context) {
	m:= "OK"
	c.JSON(200, gin.H{
		"message":m,
	})
	log.Println(m)
}
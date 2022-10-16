package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/getAllRequests", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/submitRequest", func(c *gin.Context) {

	})
	r.GET("/getAllRequestsWaitingForApproval", func(c *gin.Context) {

	})
	r.POST("/acceptRequest", func(c *gin.Context) {

	})

	r.POST("/rejectRequest", func(c *gin.Context) {

	})

	r.GET("/getAllLands", func(c *gin.Context) {

	})

	r.GET("/pay", func(c *gin.Context) {

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

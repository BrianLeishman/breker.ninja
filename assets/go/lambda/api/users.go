package main

import "github.com/gin-gonic/gin"

func init() {
	r.POST("/users", hUsersCreate)
}

func hUsersCreate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

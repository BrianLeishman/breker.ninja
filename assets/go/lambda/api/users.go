package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

func init() {
	r.POST("/users", hUsersCreate)
}

func hUsersCreate(c *gin.Context) {
	var req struct {
		Email           string `form:"email" binding:"required,email"`
		Name            string `form:"name" binding:"required,name"`
		Password        string `form:"password" binding:"required"`
		PasswordConfirm string `form:"passwordConfirm" binding:"required"`
	}
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := xid.New()

	c.JSON(200, gin.H{
		"email":  req.Email,
		"name":   req.Name,
		"userID": userID,
	})
}

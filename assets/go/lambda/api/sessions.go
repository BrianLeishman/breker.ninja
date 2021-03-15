package main

import (
	"encoding/base64"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func init() {
	r.POST("/sessions", hSessionCreate)
}

func hSessionCreate(c *gin.Context) {
	var req struct {
		Email    string `binding:"required,email"`
		Password string `binding:"required"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var users []User
	err = table.Get("sk", "email_"+req.Email).Index("GSI_1").Limit(1).All(&users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to get user").Error()})
		return
	}

	var u *User
	for i := range users {
		u = &users[i]
	}

	const passErr = "user not found or password not correct"

	if u == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": passErr})
		return
	}

	match, err := argon2id.ComparePasswordAndHash(req.Password, u.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to create password hash").Error()})
		return
	}

	if !match {
		c.JSON(http.StatusBadRequest, gin.H{"error": passErr})
		return
	}

	if !u.Verified {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not verified"})
		return
	}

	session := base64.StdEncoding.EncodeToString([]byte(u.UserID + ":" + u.Password))

	c.SetCookie("session", session, 60*60*24*365, "/", "breker.ninja", true, true)

	c.JSON(http.StatusNoContent, nil)
}

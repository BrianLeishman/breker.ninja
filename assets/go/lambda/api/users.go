package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

func init() {
	r.POST("/users", hUserCreate)
	r.PATCH("/users/:user/verify", hUserVerify)
	r.GET("/users/:user", hUserGet)
}

func hUserCreate(c *gin.Context) {
	var req struct {
		Email           string `binding:"required,email,max=255"`
		Name            string `binding:"required,name,max=255"`
		Password        string `binding:"required,min=10"`
		PasswordConfirm string `binding:"required,eqfield=Password"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: do this in one query https://stackoverflow.com/questions/55106784/how-to-insert-to-dynamodb-just-if-the-key-does-not-exist
	count, err := table.Get("sk", "email_"+req.Email).Index("GSI_1").Limit(1).Count()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to check existing email").Error()})
		return
	}

	if count != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "another user with this email address already exists"})
		return
	}

	userID := "user_" + xid.New().String()

	authKey := make([]byte, 32)
	rand.Read(authKey)

	apiKey, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to create apiKey uuid").Error()})
		return
	}

	pass, err := argon2id.CreateHash(req.Password, argon2id.DefaultParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to create password hash").Error()})
		return
	}

	err = table.Put(User{
		UserID:   userID,
		Email:    "email_" + req.Email,
		Password: pass,
		AuthKey:  authKey,
		Verified: false,
		APIKey:   apiKey.String(),
		Name:     req.Name,
		Created:  time.Now(),
	}).Run()

	c.JSON(http.StatusCreated, gin.H{
		"email": req.Email,
		"name":  req.Name,
	})
}

func hUserVerify(c *gin.Context) {
	var req struct {
		Email   string `binding:"required,email,max=255"`
		AuthKey []byte `binding:"required,len=32"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var u User
	err = table.Update("pk", "user_"+c.Param("user")).Range("sk", "email_"+req.Email).If("'authKey' = ?", req.AuthKey).Set("verified", true).Value(&u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": errors.Wrapf(err, "failed to update user's verified status").Error()})
		return
	}

	session := base64.StdEncoding.EncodeToString([]byte(u.UserID + ":" + u.Password))

	c.SetCookie("session", session, 60*60*24*365, "/", "breker.ninja", true, true)

	c.JSON(http.StatusNoContent, nil)
}

func hUserGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "brian",
	})
}

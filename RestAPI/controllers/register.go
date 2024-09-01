package controllers

import (
	"RESTAPI/initializers"
	"RESTAPI/models"
	"context"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Prepare the user data
	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}

	// Insert the user into the database
	_, err = initializers.DB.Exec(context.Background(), "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

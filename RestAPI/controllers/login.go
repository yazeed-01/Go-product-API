package controllers

import (
	"RESTAPI/initializers"
	"RESTAPI/models"
	"context"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
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

	var user models.User
	err := initializers.DB.QueryRow(context.Background(), "SELECT id, name, email, password FROM users WHERE email = $1", body.Email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to create token",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}

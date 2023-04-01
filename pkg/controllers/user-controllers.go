package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hammad-umar/goland-gin-crud-api/pkg/config"
	"github.com/hammad-umar/goland-gin-crud-api/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	// Read Body
	var body struct {
		Email 	 string 
		Password string 
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body!",
		})

		return
	}

	// Hash the password 
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password!",
		})

		return 
	}

	// Create user 
	db := config.GetDB()
	result := db.Create(&models.User{Email: body.Email, Password: string(hash)})

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is already taken!",
		})

		return 	
	}

	// Respond
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Registered Successfully!",
	})
}

func Login(ctx *gin.Context) {
	// Read Body
	var body struct {
		Email 	 string 
		Password string 
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body!",
		})

		return
	}

	// Find the user by email
	var user models.User
	db := config.GetDB()

	db.Where("email = ?", body.Email).Find(&user)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Credentials!",
		})

		return
	}

	// Compare the password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Credentials!",
		})

		return
	}

	// Generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("thisismysecret"))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate JWT token!",
		})

		return
	}

	// Set cookie
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{})
}

func Me(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

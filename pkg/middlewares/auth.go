package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hammad-umar/goland-gin-crud-api/pkg/config"
	"github.com/hammad-umar/goland-gin-crud-api/pkg/models"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

    token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("thisismysecret"), nil
    })

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		db := config.GetDB()
		db.First(&user, claims["id"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user)
		c.Next()
    } else {
		c.AbortWithStatus(http.StatusUnauthorized)
    }
}

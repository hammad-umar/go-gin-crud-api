package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hammad-umar/goland-gin-crud-api/pkg/controllers"
	"github.com/hammad-umar/goland-gin-crud-api/pkg/middlewares"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/api/users/signup", controllers.SignUp)
	r.POST("/api/users/login", controllers.Login)
	r.GET("/api/users/me", middlewares.RequireAuth,controllers.Me)
}

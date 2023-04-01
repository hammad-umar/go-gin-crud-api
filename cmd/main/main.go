package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hammad-umar/goland-gin-crud-api/pkg/routes"
)

func main() {
	r := gin.Default()

	routes.RegisterNoteRoutes(r)
	routes.RegisterUserRoutes(r)

	r.Run(":1337")
}

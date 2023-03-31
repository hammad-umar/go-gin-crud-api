package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hammad-umar/goland-gin-crud-api/pkg/controllers"
)

func RegisterNoteRoutes(r *gin.Engine) {
	r.POST("/api/notes", controllers.CreateNote)
	r.GET("/api/notes", controllers.GetAllNotes)
	r.GET("/api/notes/:id", controllers.GetSingleNote)
	r.DELETE("/api/notes/:id", controllers.DeleteNote)
	r.PATCH("/api/notes/:id", controllers.UpdateNote)
}

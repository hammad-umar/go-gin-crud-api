package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hammad-umar/goland-gin-crud-api/pkg/config"
	"github.com/hammad-umar/goland-gin-crud-api/pkg/models"
)

func GetAllNotes(ctx *gin.Context) {
	notes := models.Find()

	ctx.JSON(http.StatusOK, notes)
}

func GetSingleNote(ctx *gin.Context) {
	noteId := ctx.Param("id")
	id, _ := strconv.ParseInt(noteId, 0, 0)

	note := models.FindOneById(id)

	if note.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Note not found!",
		})

		return 
	}

	ctx.JSON(http.StatusOK, note)
}

func CreateNote(ctx *gin.Context) {
	note := models.Note{}

	ctx.BindJSON(&note)
	createdNote := note.Create()

	ctx.JSON(http.StatusCreated, createdNote)
}

func DeleteNote(ctx *gin.Context) {
	noteId := ctx.Param("id")
	id, _ := strconv.ParseInt(noteId, 0, 0)

	note := models.DeleteOneById(id)
	
	ctx.JSON(http.StatusOK, note)
}

func UpdateNote(ctx *gin.Context) {
	noteId := ctx.Param("id")
	id, _ := strconv.ParseInt(noteId, 0, 0)
	
	note := models.Note{}
	db := config.GetDB()

	ctx.BindJSON(&note)

	db.Where("id = ?", id).Updates(&models.Note{
		Title: note.Title,
		Description: note.Description,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Note is updated!",
	})
}

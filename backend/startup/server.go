package startup

import (
	"github.com/gin-gonic/gin"
	"github.com/gonotes/api/note"
)

func Server() {
	noteRepo := note.NewInMemNoteRepository()
	noteService := note.NewService(noteRepo)
	noteController := note.NewController(noteService)

	router := gin.Default()
	router.POST("/notes", noteController.PostNewNote)
	router.Run("localhost:8080")
}

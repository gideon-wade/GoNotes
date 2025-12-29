package startup

import (
	"github.com/gin-gonic/gin"
	"github.com/gonotes/api/note"
)

func Server() {
	// setup dependencies
	noteRepo := note.NewInMemNoteRepository()
	noteService := note.NewService(noteRepo)
	noteController := note.NewController(noteService)

	router := gin.Default()

	// setup routes
	api := router.Group("/api") 
	{	
		v1 := api.Group("/v1")
		{
			v1.POST("/notes", noteController.PostNewNote)
		}
	}

	router.Run("localhost:8080")
}
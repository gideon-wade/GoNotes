package startup

import (
	"github.com/gin-gonic/gin"
	"github.com/gonotes/api/note"
	"github.com/gonotes/api/logging"
)

func Server() {
	// setup dependencies
	noteRepo := note.NewInMemNoteRepository()
	noteService := note.NewService(noteRepo)
	noteControllerLogger := logging.NewStdOutLogger()
	noteController := note.NewController(noteService, noteControllerLogger)

	router := gin.Default()

	// setup routes
	api := router.Group("/api") 
	{	
		v1 := api.Group("/v1")
		{
			v1.POST("/notes", noteController.PostNewNote)
			v1.GET("/notes", noteController.GetAllNotes)
			v1.GET("/notes/:id", noteController.GetNoteByID)
		}
	}

	router.Run("localhost:8080")
}
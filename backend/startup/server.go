package startup

import (
	"io"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/gonotes/api/logging"
	"github.com/gonotes/api/note"
)

func Server() {
	// setup dependencies
	noteRepo := note.NewInMemNoteRepository()
	noteService := note.NewService(noteRepo)
	noteControllerLogger := logging.NewStdOutLogger()
	noteController := note.NewController(noteService, noteControllerLogger)

	// setup gin
	if (!gin.IsDebugging()) {
		gin.DisableConsoleColor()
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
	}

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
package startup

import (
	"github.com/gin-gonic/gin"
	"github.com/gonotes/api/auth"
	"github.com/gonotes/api/note"
)

func Server() {
	// setup dependencies
	jwtSecret := []byte("go-notes-secret-change-in-production")

	// auth
	userRepo := auth.NewInMemUserRepository()
	tokenRepo := auth.NewInMemRefreshTokenRepository()
	authService := auth.NewService(userRepo, tokenRepo, jwtSecret)
	authController := auth.NewController(authService)
	// note
	noteRepo := note.NewInMemNoteRepository()
	noteService := note.NewService(noteRepo)
	noteController := note.NewController(noteService)

	router := gin.Default()

	// setup routes
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// public endpoints
			v1.POST("/auth/register", authController.Register)
			v1.POST("/auth/login", authController.Login)
			v1.POST("/auth/refresh", authController.Refresh)
			v1.POST("/auth/logout", authController.Logout)

			// protected endpoints
			protected := v1.Group("")
			protected.Use(auth.AuthMiddleware(jwtSecret))
			{
				protected.POST("/notes", noteController.PostNewNote)
				protected.GET("/notes", noteController.GetAllNotes)
				protected.GET("/notes/:id", noteController.GetNoteByID)
			}
		}
	}

	router.Run("localhost:8080")
}

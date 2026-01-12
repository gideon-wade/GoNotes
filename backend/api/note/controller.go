package note

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gonotes/api/error"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) PostNewNote(ctx *gin.Context) {
	var newNoteRequest NewNoteRequestDTO
	err := ctx.BindJSON(&newNoteRequest)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusBadRequest,
			error.NewBadRequestError("Invalid request body."),
		)
	} else {
		newNote, err := ctrl.service.CreateNewNote(newNoteRequest)
		if err != nil {
			ctx.IndentedJSON(
				http.StatusInternalServerError,
				error.NewInternalServerError("Failed to create note."),
			)
		} else {
			ctx.IndentedJSON(http.StatusCreated, newNote)
		}
	}
}

func (ctrl *Controller) GetAllNotes(ctx *gin.Context) {
	notes, err := ctrl.service.GetAllNotes()

	if err != nil {
		ctx.IndentedJSON(
			http.StatusInternalServerError,
			error.NewInternalServerError("Failed to get notes."),
		)
	} else {
		ctx.IndentedJSON(http.StatusOK, notes)
	}
}

func (ctrl *Controller) GetNoteByID(ctx *gin.Context) {
	noteID := ctx.Param("id")

	note, err := ctrl.service.GetNoteByID(noteID)

	if err != nil {
		ctx.IndentedJSON(
			http.StatusInternalServerError,
			error.NewInternalServerError("Failed to get notes."),
		)
	} else {
		if note == nil {
			ctx.IndentedJSON(
				http.StatusNotFound,
				error.NewNotFoundError(
					fmt.Sprintf("Note with id %s not found.", noteID)),
			)
		} else {
			ctx.IndentedJSON(http.StatusOK, note)
		}
	}
}

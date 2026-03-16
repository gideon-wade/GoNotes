package note

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gonotes/api/errors"
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
			errors.NewBadRequestError("Invalid request body.", "The json body could not be serialized to the expected format.", "400-invalid-body"),
		)
	} else {
		newNote, err := ctrl.service.CreateNewNote(newNoteRequest)
		if err != nil {
			ctx.IndentedJSON(
				http.StatusInternalServerError,
				errors.NewStandardInternalServerError("Failed to create note."),
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
			errors.NewStandardInternalServerError("Failed to get notes."),
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
			errors.NewStandardInternalServerError("Failed to get notes."),
		)
	} else {
		if note == nil {
			ctx.IndentedJSON(
				http.StatusNotFound,
				errors.NewNotFoundError(
					"Note not found.",
					fmt.Sprintf("Note with id %s not found.", noteID),
					"note-not-found",
				),
			)
		} else {
			ctx.IndentedJSON(http.StatusOK, note)
		}
	}
}

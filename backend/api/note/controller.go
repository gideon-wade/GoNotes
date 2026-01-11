package note

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gonotes/api/error"
	"github.com/gonotes/api/logging"
)

type Controller struct {
	service *Service
	logger logging.Logger
}

func NewController(service *Service, logger logging.Logger) *Controller {
	return &Controller{service: service, logger: logger}
}

func (ctrl *Controller) PostNewNote(ctx *gin.Context) {
	ctrl.logger.Log(logging.RequestReceived(ctx))
	var newNoteRequest NewNoteRequestDTO
	err := ctx.BindJSON(&newNoteRequest)
	if err != nil {
		ctrl.logger.Log(logging.NewLogEventError(
			"Invalid request body.",
			err,	
		))
		ctx.IndentedJSON(
			http.StatusBadRequest,
			error.NewBadRequestError("Invalid request body."),
		)
	} else {
		newNote, err := ctrl.service.CreateNewNote(newNoteRequest)
		if err != nil {
			ctrl.logger.Log(logging.NewLogEventError(
				"Failed to create note.",
				err,
			))
			ctx.IndentedJSON(
				http.StatusInternalServerError,
				error.NewInternalServerError("Failed to create note."),
			)
		} else {
			ctrl.logger.Log(logging.RequestCompleted(ctx))
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

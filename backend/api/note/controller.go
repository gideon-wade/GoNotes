package note

import (
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
			errors.NewBadRequestError("Invalid request body.", "The json body could not be serialized to the expected format."),
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

package note

import (
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
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to create note"})
		} else {
			ctx.IndentedJSON(http.StatusCreated, newNote)
		}
	}

}

package note

import (
	"net/http"
	"github.com/gin-gonic/gin"
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
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": "invalid note",
		})
	} else {
		newNote, err := ctrl.service.CreateNewNote(newNoteRequest)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to create note"})
		} else {
			ctx.IndentedJSON(http.StatusCreated, newNote)
		}
	}

}

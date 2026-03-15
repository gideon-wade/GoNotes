package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	apiError "github.com/gonotes/api/error"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func (ctrl *Controller) Register(ctx *gin.Context) {
	var request RegisterRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		errResp := apiError.NewBadRequestError("Invalid request body.")
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}

	user, errResp := ctrl.service.Register(request)
	if errResp != nil {
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}

	ctx.IndentedJSON(http.StatusCreated, user)
}

func (ctrl *Controller) Login(ctx *gin.Context) {
	var request LoginRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		errResp := apiError.NewBadRequestError("Invalid request body.")
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}

	authResponse, errResp := ctrl.service.Login(request)
	if errResp != nil {
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}

	ctx.IndentedJSON(http.StatusOK, authResponse)
}

func (ctrl *Controller) Refresh(ctx *gin.Context) {
	var request RefreshTokenRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		errResp := apiError.NewBadRequestError("Invalid request body.")
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}

	authResponse, errResp := ctrl.service.Refresh(request)
	if errResp != nil {
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}

	ctx.IndentedJSON(http.StatusOK, authResponse)
}

func (ctrl *Controller) Logout(ctx *gin.Context) {
	var request RefreshTokenRequestDTO
	if err := ctx.BindJSON(&request); err != nil {
		errResp := apiError.NewBadRequestError("Invalid request body.")
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}

	errResp := ctrl.service.Logout(request)
	if errResp != nil {
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}

	ctx.Status(http.StatusNoContent)
}

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
	var registerRequest RegisterRequestDTO
	err := ctx.BindJSON(&registerRequest)
	if err != nil {
		errResp := apiError.NewBadRequestError("Invalid request body.")
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}

	user, errResp := ctrl.service.Register(registerRequest)
	if errResp != nil {
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}
	ctx.IndentedJSON(http.StatusCreated, user)
}

func (ctrl *Controller) Login(ctx *gin.Context) {
	var loginRequest LoginRequestDTO
	err := ctx.BindJSON(&loginRequest)
	if err != nil {
		errResp := apiError.NewBadRequestError("Invalid request body.")
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}

	authResponse, errResp := ctrl.service.Login(loginRequest)
	if errResp != nil {
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}
	ctx.IndentedJSON(http.StatusOK, authResponse)
}

func (ctrl *Controller) Refresh(ctx *gin.Context) {
	var refreshRequest RefreshTokenRequestDTO
	err := ctx.BindJSON(&refreshRequest)
	if err != nil {
		errResp := apiError.NewBadRequestError("Invalid request body.")
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}

	authResponse, errResp := ctrl.service.Refresh(refreshRequest)
	if errResp != nil {
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}
	ctx.IndentedJSON(http.StatusOK, authResponse)
}

func (ctrl *Controller) Logout(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	if userID == "" {
		errResp := apiError.NewUnauthorizedError("Unauthorized.")
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}

	errResp := ctrl.service.Logout(userID)
	if errResp != nil {
		ctx.IndentedJSON(errResp.Status, errResp)
		return
	}
	ctx.Status(http.StatusNoContent)
}

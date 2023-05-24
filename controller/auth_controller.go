package controller

import (
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/test2/data/request"
	"github.com/sferawann/test2/data/response"
	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/service"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{authService: service}
}

func (c *AuthController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	helper.ErrorPanic(err)

	token, err_token := c.authService.Login(loginRequest)
	fmt.Println(err_token)
	if err_token != nil {
		webResponse := response.Response{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: "Invalid username or password",
		}
		ctx.JSON(http.StatusBadRequest, webResponse)
		return
	}

	resp := response.LoginResponse{
		TokenType: "Bearer",
		Token:     token,
	}

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully log in!",
		Data:    resp,
	}

	// ctx.SetCookie("token", token, config.TokenMaxAge*60, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, webResponse)
}

func (c *AuthController) Register(ctx *gin.Context) {
	createBorrowerRequest := request.CreateBorrowerRequest{}
	err := ctx.ShouldBindJSON(&createBorrowerRequest)
	helper.ErrorPanic(err)

	c.authService.Register(createBorrowerRequest)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created user!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

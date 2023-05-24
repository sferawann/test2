package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sferawann/test2/data/response"
	"github.com/sferawann/test2/helper"
	"github.com/sferawann/test2/model"
	"github.com/sferawann/test2/service"
)

type BorrowerController struct {
	borrowerService service.BorrowerService
}

func NewBorrowerController(service service.BorrowerService) *BorrowerController {
	return &BorrowerController{borrowerService: service}
}

func (c *BorrowerController) Insert(ctx *gin.Context) {
	var bor model.Borrower

	err := ctx.ShouldBindJSON(&bor)
	helper.ErrorPanic(err)

	c.borrowerService.Save(bor)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully Add data Borrower!",
		Data:    bor,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *BorrowerController) Update(ctx *gin.Context) {
	panic("unimplemented")
}
func (c *BorrowerController) Delete(ctx *gin.Context) {
	panic("unimplemented")
}

func (c *BorrowerController) FindAll(ctx *gin.Context) {
	bor := c.borrowerService.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all borrower data!",
		Data:    bor,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *BorrowerController) FindByID(ctx *gin.Context) {
	panic("unimplemented")
}

func (c *BorrowerController) FindByUsername(ctx *gin.Context) {
	panic("unimplemented")
}

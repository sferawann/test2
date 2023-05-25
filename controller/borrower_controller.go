package controller

import (
	"net/http"
	"strconv"

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
	createBor := model.Borrower{}
	err := ctx.ShouldBindJSON(&createBor)
	helper.ErrorPanic(err)

	c.borrowerService.Save(createBor)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created Borrower!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *BorrowerController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	updateBor := model.Borrower{}
	err = ctx.ShouldBindJSON(&updateBor)
	helper.ErrorPanic(err)

	updateBor.Id = id
	c.borrowerService.Update(updateBor)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated Borrower!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *BorrowerController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	c.borrowerService.Delete(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted Borrower!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
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
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	bor, err := c.borrowerService.FindById(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Borrower!",
		Data:    bor,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *BorrowerController) FindByUsername(ctx *gin.Context) {
	userParam := ctx.Param("username")

	bor, err := c.borrowerService.FindByUsername(userParam)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Borrower!",
		Data:    bor,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

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

type LoanProductController struct {
	loanProductService service.LoanProductService
}

func NewLoanProductController(service service.LoanProductService) *LoanProductController {
	return &LoanProductController{loanProductService: service}
}

func (c *LoanProductController) Insert(ctx *gin.Context) {
	createLP := model.LoanProduct{}
	err := ctx.ShouldBindJSON(&createLP)
	helper.ErrorPanic(err)

	c.loanProductService.Save(createLP)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created Lender!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *LoanProductController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	updateLP := model.LoanProduct{Id: id}
	err = ctx.ShouldBindJSON(&updateLP)
	helper.ErrorPanic(err)

	updatedLender, err := c.loanProductService.Update(updateLP)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated Lender!",
		Data:    updatedLender,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *LoanProductController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	c.loanProductService.Delete(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted Lender!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *LoanProductController) FindAll(ctx *gin.Context) {
	lp, err := c.loanProductService.FindAll()
	helper.ErrorPanic(err)
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all Lender data!",
		Data:    lp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *LoanProductController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	lp, err := c.loanProductService.FindById(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Lender!",
		Data:    lp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *LoanProductController) FindByName(ctx *gin.Context) {
	userParam := ctx.Query("name")

	lp, err := c.loanProductService.FindByName(userParam)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Lender!",
		Data:    lp,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

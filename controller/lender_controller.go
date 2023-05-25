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

type LenderController struct {
	lenderService service.LenderService
}

func NewLenderController(service service.LenderService) *LenderController {
	return &LenderController{lenderService: service}
}

func (c *LenderController) Insert(ctx *gin.Context) {
	createLen := model.Lender{}
	err := ctx.ShouldBindJSON(&createLen)
	helper.ErrorPanic(err)

	c.lenderService.Save(createLen)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created Lender!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *LenderController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	updateLen := model.Lender{}
	err = ctx.ShouldBindJSON(&updateLen)
	helper.ErrorPanic(err)

	updateLen.Id = id
	c.lenderService.Update(updateLen)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully updated Lender!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
func (c *LenderController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	c.lenderService.Delete(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully deleted Lender!",
		Data:    nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *LenderController) FindAll(ctx *gin.Context) {
	bor := c.lenderService.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetch all Lender data!",
		Data:    bor,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (c *LenderController) FindByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	helper.ErrorPanic(err)

	bor, err := c.lenderService.FindById(id)
	helper.ErrorPanic(err)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully fetched Lender!",
		Data:    bor,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
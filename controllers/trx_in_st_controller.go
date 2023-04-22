package controllers

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrxInStController struct {
	usecase usecase.TrxInStUsecase
}

func (c *TrxInStController) EnrollInsertTrxInSt(ctx *gin.Context) {
	var newtrxInSt entity.TrxInST

	if err := ctx.BindJSON(&newtrxInSt); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	res := c.usecase.EnrollInsertTrx(&newtrxInSt)

	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func NewTrxInStController(u usecase.TrxInStUsecase) *TrxInStController {
	controller := TrxInStController{
		usecase: u,
	}

	return &controller
}

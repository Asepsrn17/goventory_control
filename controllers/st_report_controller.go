package controllers

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportTrxStController struct {
	usecase usecase.ReportTrxStUsecase
}

func (c *ReportTrxStController) FindAllReportrxSt(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// username := claims["username"].(string)
	res := c.usecase.FindAllReportTrxSt()
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (c *ReportTrxStController) FindByReportTrxProductStId(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// username := claims["username"].(string)
	var reportsTrxSt entity.ReportTrxSt
	if err := ctx.BindJSON(&reportsTrxSt); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}
	res := c.usecase.FindByReportTrxProductStId(reportsTrxSt.ProductStId)
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func (c *ReportTrxStController) FindByDateReportTrxSt(ctx *gin.Context) {
	// claims := ctx.MustGet("claims").(jwt.MapClaims)
	// username := claims["username"].(string)

	var reportsTrxSt entity.ReportTrxSt
	if err := ctx.BindJSON(&reportsTrxSt); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}
	res := c.usecase.FindByDateReportTrxSt(reportsTrxSt.CreatedAt)
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

func NewReportTrxStController(u usecase.ReportTrxStUsecase) *ReportTrxStController {
	controller := ReportTrxStController{
		usecase: u,
	}

	return &controller
}

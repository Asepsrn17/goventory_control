package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	"github.com/gin-gonic/gin"
)

func TrxStRoutes(router *gin.Engine, db *sql.DB) {

	//controller Trx ProductSt

	trxInstRepo := repository.NewTrxInStRepo(db)
	trxInStUsecase := usecase.NewTrxInStUsecase(trxInstRepo)
	trxStctrl := controllers.NewTrxInStController(trxInStUsecase)

	//cotroller Report Trx Product St
	reportTrxStRepo := repository.NewReportTrxStRepo(db)
	reportTrxStUsecase := usecase.NewReportTrxStUsecase(reportTrxStRepo)
	reportTrxStctrl := controllers.NewReportTrxStController(reportTrxStUsecase)

	//Router ProductSt dan TrxSt
	trxInStRouter := router.Group("/api/v1/trxinstore")
	trxInStRouter.POST("", trxStctrl.EnrollInsertTrxInSt)
	trxInStRouter.GET("/report_all", reportTrxStctrl.FindAllReportrxSt)
	trxInStRouter.GET("/reportby_st_id", reportTrxStctrl.FindByReportTrxProductStId)
	trxInStRouter.GET("/reportby_st_date", reportTrxStctrl.FindByDateReportTrxSt)
}

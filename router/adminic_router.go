package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/middleware"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	"github.com/gin-gonic/gin"
)

func RouterAdminIc(router *gin.Engine, db *sql.DB) {
	adminIcRepo := repository.NewAdminIcRepo(db)
	adminIcUsecase := usecase.NewAdminIcUsecase(adminIcRepo)
	adminCtrl := controllers.NewAdminIcController(adminIcUsecase)

	router.POST("/auth/login/adminic", middleware.Login)

	adminRouter := router.Group("/icteam")

	adminRouter.GET("/:id/profile/adminic", middleware.AuthMiddleware(), middleware.Profile)
	adminRouter.GET("", adminCtrl.FindAdminIc)
	adminRouter.GET("/:id", adminCtrl.FindAdminIcByid)
	adminRouter.POST("", adminCtrl.Register)
	adminRouter.PUT("", adminCtrl.Edit)
	adminRouter.DELETE("/:id", adminCtrl.Unreg)
}
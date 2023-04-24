package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	"github.com/gin-gonic/gin"
)

func InitRouterEmployee(router *gin.Engine, db *sql.DB) {
	// Dependencies Warehouse Team
	warehouseTeamRepo := repository.NewWarehouseTeamRepo(db)
	warehouseTeamUsecase := usecase.NewWarehouseTeamUsecase(warehouseTeamRepo)
	warehouseTeamController := controllers.NewWarehouseTeamController(warehouseTeamUsecase)

	//Register and Login session
	router.POST("/register", warehouseTeamController.Register)
	router.POST("/auth/login", warehouseTeamController.Login)

	// group users
	userRouter := router.Group("/warehouse-team/employees")

	// routes (GET, POST, PUT, DELETE)
	userRouter.GET("", warehouseTeamController.FindEmployees)
	userRouter.GET("/:id", warehouseTeamController.FindEmployeeById)
	userRouter.PUT("", warehouseTeamController.Edit)
	userRouter.DELETE("/:id", warehouseTeamController.Unreg)

}

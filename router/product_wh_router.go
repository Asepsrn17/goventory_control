package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	"github.com/gin-gonic/gin"
)

func InitRouterProduct(router *gin.Engine, db *sql.DB) {
	// Dependencies Product Warehouse
	productWhRepo := repository.NewProductWhRepo(db)
	productWhUsecase := usecase.NewProductWhUsecase(productWhRepo)
	productWhController := controllers.NewProductWhController(productWhUsecase)

	// group users
	userRouter := router.Group("/warehouse/products")

	// routes (GET, POST, PUT, DELETE)
	userRouter.GET("", productWhController.FindProducts)
	userRouter.GET("/:id", productWhController.FindProductById)
	userRouter.GET("/name/:name", productWhController.FindProductByName)
	userRouter.POST("/input", productWhController.Input)
	userRouter.PUT("", productWhController.Edit)
	userRouter.DELETE("/:id", productWhController.Output)
}

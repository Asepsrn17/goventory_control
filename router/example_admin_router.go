package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	// "github.com/Uchel/auth-jwt2/middleware"
	jwt_controller "github.com/Uchel/auth-jwt2/controller"
	jwt_middleware "github.com/Uchel/auth-jwt2/middleware"
	jwt_repository "github.com/Uchel/auth-jwt2/repository"
	jwt_usecase "github.com/Uchel/auth-jwt2/usecase"
	"github.com/gin-gonic/gin"
)

func LoginExampleRoutes(router *gin.Engine, db *sql.DB) {
	// auth jwt login ExampleAdmin
	jwtExampleAdminRepo := jwt_repository.NewStTeamLoginRepo(db)
	jwtExampleAdminUsecase := jwt_usecase.NewStTeamLoginUsecase(jwtExampleAdminRepo)
	jwtExampleAdminCtrl := jwt_controller.NewStTeamLoginController(jwtExampleAdminUsecase)

	//example Admin
	exampleAdminRepo := repository.NewExampleAdminRepo(db)
	exampleAdminUsecase := usecase.NewExampleAdminUsecase(exampleAdminRepo)
	exampleAdminctrl := controllers.NewExampleAdminController(exampleAdminUsecase)

	//Register session
	userRegister := router.Group("/register")
	userRegister.POST("/example-admin", exampleAdminctrl.RegisterExampleAdmin)

	// login sesion
	// router.POST("/auth/login", jwtExampleAdminCtrl.LoginStTeam)
	loginRouteGroup := router.Group("/auth/login")
	loginRouteGroup.POST("/example-admin", jwtExampleAdminCtrl.LoginStTeam)

	//Jwt Access
	exampleadminGroup := router.Group("/example_admin_profil")
	exampleadminGroup.Use(jwt_middleware.AuthMiddleware())
	exampleadminGroup.GET("/get_all", exampleAdminctrl.FindAllExampleAdmin)
	exampleadminGroup.GET("/get_by_id", exampleAdminctrl.FindByIdExampleAdmin)
	exampleadminGroup.PUT("/update", exampleAdminctrl.EditExampleAdmin)
	exampleadminGroup.DELETE("delete", exampleAdminctrl.UnregExampleAdmin)

}

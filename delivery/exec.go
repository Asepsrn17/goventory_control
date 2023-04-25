package delivery

import (
	"go_inven_ctrl/config"
	"go_inven_ctrl/router"

	"github.com/gin-gonic/gin"
)

func Exec() {
	r := gin.Default()
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}
	// router.LoginExampleRoutes(r, db)
	// router.ProducStRoutes(r, db)
	// router.TrxStRoutes(r, db)
	//router.InitRouterEmployee(r, db)
	router.InitRouterProduct(r, db)
	router.RouterAdminIc(r, db)

	r.Run(":8080")
}

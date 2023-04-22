package delivery

import (
	"go_inven_ctrl/config"
	"go_inven_ctrl/router"

	"github.com/gin-gonic/gin"
)

func Exec() {
	r := gin.Default()
	db := config.ConnectDB()

	router.LoginExampleRoutes(r, db)
	router.ProducStRoutes(r, db)
	router.TrxStRoutes(r, db)

	r.Run(":7000")
}

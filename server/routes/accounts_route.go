package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vyas-git/demyst-reports/controllers/accounts"
)

func initAccountsRoutes(rg *gin.RouterGroup) {

	adminRoute := rg.Group("/accounts")
	adminRoute.GET("/reports/balance-sheet", accounts.GetBalanceSheet)

}

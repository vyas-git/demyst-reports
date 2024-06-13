package accounts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vyas-git/demyst-reports/log"
)

func GetBalanceSheet(ctx *gin.Context) {
	var params struct {
	}

	if err := ctx.Bind(&params); err != nil {
		log.LogMessage(
			"balance sheet",
			"invalid parameter to get balance sheet",
			"error",
			logrus.Fields{
				"error": err.Error(),
				"uri":   "/api/accounts/reports/balance-sheet",
			},
		)
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			err.Error(),
		)
		return
	}
	resp, err := reportsController.GetBalanceSheet()

	if err != nil {
		log.LogMessage("balance sheet", "unable to get response from xero", "error", logrus.Fields{
			"error": err.Error(),
			"uri":   "/api/accounts/reports/balance-sheet",
		})
	}
	ctx.JSON(http.StatusOK, resp)
}

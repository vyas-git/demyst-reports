package accounts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vyas-git/demyst-reports/config"
	"github.com/vyas-git/demyst-reports/log"
)

type ReportsController struct {
}

func (c *ReportsController) Init() {
	log.LogMessage(
		"accounts_reports_controller",
		"successfully initialized",
		"success",
		logrus.Fields{},
	)
}

var reportsController ReportsController

func InitReportsController() {
	reportsController.Init()
}

func (c *ReportsController) GetBalanceSheet() (*BalanceSheet, error) {
	config := config.Get()
	xero_api := config.XeroAPI

	balance_sheet_api := fmt.Sprintf("%v/Reports/BalanceSheet", xero_api)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(balance_sheet_api)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var balanceSheet BalanceSheet
	err = json.Unmarshal(body, &balanceSheet)
	if err != nil {
		return nil, err
	}

	return &balanceSheet, nil
}

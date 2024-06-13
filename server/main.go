package main

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vyas-git/demyst-reports/config"
	"github.com/vyas-git/demyst-reports/controllers/accounts"
	"github.com/vyas-git/demyst-reports/log"
	"github.com/vyas-git/demyst-reports/routes"
)

func initConfig() *config.Config {
	config.Init()
	config := config.Get()
	// log.LogMessage("main thread", "initializing config done...", "success", logrus.Fields{})
	return &config
}

func initLog() {
	log.Init()
	log.LogMessage("main thread", "initializing log done...", "success", logrus.Fields{})
}

func initRoute(config *config.Config) {
	r := routes.Init()
	if config.ENV == "dev" {
		r.Use(static.Serve("/", static.LocalFile("../build", true)))
		r.NoRoute(func(c *gin.Context) {
			c.File("../build/index.html")
		})
	}

	if err := r.Run(fmt.Sprintf(":%v", config.AppPort)); err != nil {
		log.LogMessage("main thread", fmt.Sprintf("failed to run server %v", err), "error", logrus.Fields{})
		return
	}

	log.LogMessage("main thread", "initializing router done...", "success", logrus.Fields{})
}

func initialize() {
	// log.LogMessage("main thread", "initializing...", "info", logrus.Fields{})
	config := initConfig()
	initLog()

	accounts.InitReportsController()
	initRoute(config)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.LogMessage(
				"main thread",
				"recovered",
				"info",
				logrus.Fields{
					"recover": r,
					"stack":   string(debug.Stack()),
				},
			)
		}
	}()

	initialize()
}

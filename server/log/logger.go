package log

import (
	"encoding/json"
	"fmt"

	"github.com/TwiN/go-color"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func Init() {
	lgr := logrus.New()
	Logger = lgr
}
func LogMessage(caller string, message string, level string, fields logrus.Fields) {
	jsonString, _ := json.Marshal(fields)
	switch level {
	case "error":
		fmt.Println(color.Colorize(color.Red, "-- "+caller+" -> "+message+" -- "+string(jsonString)))
		// Logger.WithFields(fields).Errorln(caller + " -> " + message + " -- " + string(jsonString))
	case "info":
		fmt.Println(color.Colorize(color.Cyan, "-- "+caller+" -> "+message+" -- "+string(jsonString)))
		// Logger.WithFields(fields).Infoln(caller + " -> " + message + " -- " + string(jsonString))
	case "success":
		fmt.Println(color.Colorize(color.Green, "-- "+caller+" -> "+message+" -- "+string(jsonString)))
		// Logger.WithFields(fields).Println(caller + " -> " + message + " -- " + string(jsonString))
	default:
		fmt.Println(color.Colorize(color.Purple, "-- "+caller+" -> "+message+" : "+level))
		break
	}
}

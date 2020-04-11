package mylogger

import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/t-tomalak/logrus-easy-formatter"
)

var log = logrus.New()

func init() {
	log.SetOutput(os.Stderr)
	log.SetLevel(logrus.InfoLevel)

	//log.SetFormatter(&logrus.JSONFormatter{})
	//log.SetFormatter(&logrus.TextFormatter{
	//	DisableColors: true,
	//	DisableTimestamp : true,
	//})
	log.SetFormatter(
		&easy.Formatter{
		  LogFormat: "[%lvl%] %msg%\n",
	})
}

func GetLogger() *logrus.Entry {
	return log.WithFields(logrus.Fields{
		"common": "this is a common field",
		//"other": "I also should be logged always",
	})
}

func SetLevelDebug() {
	log.SetLevel(logrus.DebugLevel)
}
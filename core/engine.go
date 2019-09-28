package core

import (
	"io"
	"log"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/josephmathew1996/trendz-holidays/config"
)

// SetUpEngine initializes the app & returns a gin Handler if there is no error
func SetUpEngine(conf *config.Vars) (*gin.Engine, error) {
	engine := gin.Default()
	gin.SetMode(conf.Mode)
	setLogger(conf)
	return engine, nil
}

//setLogger initializes logrus logger with some defaults
func setLogger(conf *config.Vars) {
	file, err := os.OpenFile(conf.Log.Location, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("Error opening log file : %s\n", err)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{})
	mw := io.MultiWriter(os.Stdout, file)
	logrus.SetOutput(mw)
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}
}

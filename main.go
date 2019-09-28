package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/josephmathew1996/trendz-holidays/config"
	"github.com/josephmathew1996/trendz-holidays/core"
	"github.com/josephmathew1996/trendz-holidays/routes"
)

func main() {
	conf := config.ReadConfig()
	engine, err := core.SetUpEngine(conf)
	if err != nil {
		logrus.Fatalf("Error initialising gin engine %s\n", err)
	} else {
		logrus.Info("Gin engine initailisation success")
	}
	core.SetUpService(conf, engine)
	routes.SetUpRouter(engine)
	err = engine.Run(":" + conf.Port)
	if err != nil {
		logrus.Fatalf("Error running server %s\n", err)
	}
}

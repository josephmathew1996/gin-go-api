package core

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //import mysql
	"github.com/jmoiron/sqlx"
	"github.com/josephmathew1996/trendz-holidays/config"
)

func SetUpService(conf *config.Vars, c *gin.Engine) {
	c.Use(mysqlReadconfig(conf))
	c.Use(mysqlWriteconfig(conf))
	c.Use(loadConfig(conf))
}

func mysqlReadconfig(conf *config.Vars) gin.HandlerFunc {
	dbr := sqlx.MustConnect("mysql", conf.Mysql.Reader)
	if dbr == nil {
		logrus.Fatal("Unbale to connect to mysql")
	} else {
		logrus.Info("MySql reader connection successfull")
	}
	return func(c *gin.Context) {
		c.Set("dbr", dbr)
		c.Next()
	}
}

func mysqlWriteconfig(conf *config.Vars) gin.HandlerFunc {
	dbr := sqlx.MustConnect("mysql", conf.Mysql.Writer)
	if dbr == nil {
		logrus.Fatal("Unbale to connect to mysql")
	} else {
		logrus.Info("MySql writer connection successfull")
	}
	return func(c *gin.Context) {
		c.Set("dbw", dbr)
		c.Next()
	}
}
func loadConfig(conf *config.Vars) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", conf)
		c.Next()
	}
}

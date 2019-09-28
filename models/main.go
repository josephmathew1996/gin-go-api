package models

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/josephmathew1996/trendz-holidays/config"
)

type Service struct {
	Global *gin.Context
}

func (c *Service) GetReadDB() *sqlx.DB {

	dbr := c.Global.MustGet("dbr").(*sqlx.DB)
	return dbr

}

func (c *Service) GetWriteDB() *sqlx.DB {
	dbw := c.Global.MustGet("dbw").(*sqlx.DB)
	return dbw

}

//LoadConfig loads config values
func (c *Service) LoadConfig() *config.Vars {
	config := c.Global.MustGet("config").(*config.Vars)
	return config
}

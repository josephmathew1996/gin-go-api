package config

import (
	"log"

	envconfig "github.com/hifx/envconfig"
)

var siteConfig *Vars

// Mode Values
const (
	LOG_FORMAT_TXT  string = "txt"
	LOG_FORMAT_JSON string = "json"
	APP_NAME        string = "TRENDZ_HOLIDAYS"
)

// SiteConfig defines the site url and name details
type SiteConfig struct {
	SiteURL  string `envconfig:"TRENDZ_HOLIDAYS_SITE_URL"`
	SiteName string `envconfig:"TRENDZ_HOLIDAYS_SITE_NAME"`
}

// Vars represents the configuration for an App
type Vars struct {
	App  string `envconfig:"-"`
	Mode string `envconfig:"TRENDZ_HOLIDAYS_MODE"`
	Port string `envconfig:"TRENDZ_HOLIDAYS_PORT,default=:8080"`
	Log  struct {
		Location string `envconfig:"TRENDZ_HOLIDAYS_ERROR_LOG"`
	}
	Mysql struct {
		Reader    string `envconfig:"TRENDZ_HOLIDAYS_MYSQL_READ_DSN"`
		Writer    string `envconfig:"TRENDZ_HOLIDAYS_MYSQL_WRITE_DSN"`
		MaxActive int    `envconfig:"TRENDZ_HOLIDAYS_MYSQL_MAX_ACTIVE"`
		MaxIdle   int    `envconfig:"TRENDZ_HOLIDAYS_MYSQL_MAX_IDLE"`
	}
	Site SiteConfig
}

//ReadConfig reads enc config
func ReadConfig() *Vars {

	vars := &Vars{App: APP_NAME}
	if err := envconfig.Init(vars); err != nil {
		log.Fatalf("Error reading configuration : %s\n", err)
	}
	siteConfig = vars
	return vars
}

//GetSiteConfig returns siteconfig
func GetSiteConfig() *Vars {
	return siteConfig
}

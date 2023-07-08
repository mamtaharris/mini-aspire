package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Name      string `json:"APP_NAME"`
	Port      int    `json:"APP_PORT"`
	JwtSecret string `json:"APP_JWT_SECRET"`
}

var App *AppConfig

func loadAppConfig() {
	App = &AppConfig{}
	err := envconfig.Process("app", App)
	if err != nil {
		log.Fatal(err.Error())
	}
}

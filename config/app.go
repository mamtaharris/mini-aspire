package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	Name string `envconfig:"APP_NAME"`
	Port int    `envconfig:"APP_PORT"`
}

var App *AppConfig

func loadAppConfig() {
	App = &AppConfig{}
	err := envconfig.Process("app", App)
	if err != nil {
		log.Fatal(err.Error())
	}
}

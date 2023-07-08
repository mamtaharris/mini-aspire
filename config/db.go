package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	Host               string        `envconfig:"DB_HOST"`
	Username           string        `envconfig:"DB_USERNAME"`
	Password           string        `envconfig:"DB_PASSWORD"`
	Name               string        `envconfig:"DB_NAME"`
	Port               int           `envconfig:"DB_PORT"`
	MaxIdleConnections int           `envconfig:"DB_MAX_IDLE_CONNECTIONS"`
	MaxOpenConnections int           `envconfig:"DB_MAX_OPEN_CONNECTIONS"`
	ConnectionLifetime time.Duration `envconfig:"DB_CONNECTION_LIFETIME"`
	QueryTimeout       time.Duration `envconfig:"DB_QUERY_TIMEOUT"`
}

var DB *DBConfig

func loadDBConfig() {
	DB = &DBConfig{}
	err := envconfig.Process("db", DB)
	if err != nil {
		log.Fatal(err.Error())
	}
}

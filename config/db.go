package config

import (
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	Host               string        `json:"DB_HOST"`
	Username           string        `json:"DB_USERNAME"`
	Password           string        `json:"DB_PASSWORD"`
	Name               string        `json:"DB_NAME"`
	Port               int           `json:"DB_PORT"`
	MaxIdleConnections int           `json:"DB_MAX_IDLE_CONNECTIONS"`
	MaxOpenConnections int           `json:"DB_MAX_OPEN_CONNECTIONS"`
	ConnectionLifetime time.Duration `json:"DB_CONNECTION_LIFETIME"`
	QueryTimeout       time.Duration `json:"DB_QUERY_TIMEOUT"`
}

var DB *DBConfig

func loadDBConfig() {
	DB = &DBConfig{}
	err := envconfig.Process("db", DB)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (dc *DBConfig) ConnectionURL() string {
	if dc == nil {
		return ""
	}

	host := dc.Host
	if v := dc.Port; v > 0 {
		host = host + ":" + strconv.Itoa(v)
	}

	u := &url.URL{
		Scheme: "postgres",
		Host:   host,
		Path:   dc.Name,
	}

	if dc.Username != "" || dc.Password != "" {
		u.User = url.UserPassword(dc.Username, dc.Password)
	}

	q := u.Query()
	q.Add("sslmode", "disable")
	u.RawQuery = q.Encode()

	return u.String()
}

package config

import (
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/mamtaharris/mini-aspire/pkg/logger"
)

func InitConfig() {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../")
	err := godotenv.Load(projectRoot + "/.env")
	if err != nil {
		logger.Log.Fatal(err.Error())
	}
	loadAppConfig()
	loadDBConfig()
}

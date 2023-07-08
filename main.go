package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/mamtaharris/mini-aspire/cmd"
	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/pkg/database"
	"github.com/mamtaharris/mini-aspire/pkg/logger"
)

func main() {
	config.InitConfig()
	logger.InitLogger()
	db := database.InitDB()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	defer stop()

	err := cmd.Execute(ctx, db)
	if err != nil {
		logger.Log.Fatal(err.Error())
	}
}

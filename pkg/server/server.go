package server

import (
	"context"
	"fmt"

	"github.com/mamtaharris/mini-aspire/config"
	"github.com/mamtaharris/mini-aspire/internal/router"
	"gorm.io/gorm"
)

func Start(ctx context.Context, db *gorm.DB) error {
	router, err := router.SetRouter(ctx, db)
	if err != nil {
		return err
	}
	router.Run(":" + fmt.Sprintf("%d", config.App.Port))
	return nil
}

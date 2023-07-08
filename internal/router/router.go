package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetRouter(ctx context.Context, db *gorm.DB) (*gin.Engine, error) {
	router := gin.Default()
	router.HandleMethodNotAllowed = true
	return router, nil
}

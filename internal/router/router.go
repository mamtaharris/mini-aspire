package router

import (
	"context"

	"github.com/gin-gonic/gin"
	handler "github.com/mamtaharris/mini-aspire/internal/handlers"
	loanR "github.com/mamtaharris/mini-aspire/internal/repositories/loan"
	loanS "github.com/mamtaharris/mini-aspire/internal/services/loan"
	loanV "github.com/mamtaharris/mini-aspire/internal/validators/loan"

	"gorm.io/gorm"
)

func SetRouter(ctx context.Context, db *gorm.DB) (*gin.Engine, error) {
	router := gin.Default()
	router.HandleMethodNotAllowed = true

	loanRepo := loanR.NewRepo(db, db)
	loanSvc := loanS.NewService(loanRepo)
	loanValidator := loanV.NewValidator()
	loanHandler := handler.NewLoanHandler(loanSvc, loanValidator)

	routerV1 := router.Group("/v1")
	routerV1.POST("/loan", loanHandler.CreateLoanHandler)

	return router, nil
}

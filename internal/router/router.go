package router

import (
	"context"

	"github.com/gin-gonic/gin"
	handler "github.com/mamtaharris/mini-aspire/internal/handlers"
	"github.com/mamtaharris/mini-aspire/internal/repositories"

	"github.com/mamtaharris/mini-aspire/internal/services"

	"github.com/mamtaharris/mini-aspire/internal/validators"
	"github.com/mamtaharris/mini-aspire/pkg/middleware"

	"gorm.io/gorm"
)

func SetRouter(ctx context.Context, db *gorm.DB) (*gin.Engine, error) {
	router := gin.Default()
	router.HandleMethodNotAllowed = true

	loanRepo := repositories.NewLoanRepo(db, db)
	repaymentRepo := repositories.NewRepaymentRepo(db, db)
	userRepo := repositories.NewUserRepo(db, db)
	loanSvc := services.NewLoanService(loanRepo, repaymentRepo)
	userSvc := services.NewUserService(userRepo)
	loanValidator := validators.NewLoanValidator()
	userValidator := validators.NewUserValidator()
	loanHandler := handler.NewLoanHandler(loanSvc, loanValidator)
	userHandler := handler.NewUserHandler(userSvc, userValidator)
	authMiddleware := middleware.NewAuthMiddleware(userRepo)

	routerV1 := router.Group("/v1")
	routerV1.POST("/login", userHandler.Login)
	routerV1.POST("/loan", authMiddleware.Authenticate, authMiddleware.Authorize("ADMIN"), loanHandler.CreateLoanHandler)
	routerV1.PUT("/loan/:id", loanHandler.UpdateLoanHandler)
	routerV1.GET("/loan/:id", loanHandler.GetLoanHandler)
	routerV1.POST("/loan/:loanID/repay/:repaymentID", loanHandler.RepayLoanHandler)

	return router, nil
}

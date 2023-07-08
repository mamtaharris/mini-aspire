package router

import (
	"context"

	"github.com/gin-gonic/gin"
	handler "github.com/mamtaharris/mini-aspire/internal/handlers"
	loanR "github.com/mamtaharris/mini-aspire/internal/repositories/loan"
	repaymentR "github.com/mamtaharris/mini-aspire/internal/repositories/repayment"
	userR "github.com/mamtaharris/mini-aspire/internal/repositories/users"
	loanS "github.com/mamtaharris/mini-aspire/internal/services/loan"
	userS "github.com/mamtaharris/mini-aspire/internal/services/users"
	loanV "github.com/mamtaharris/mini-aspire/internal/validators/loan"
	userV "github.com/mamtaharris/mini-aspire/internal/validators/user"
	"github.com/mamtaharris/mini-aspire/pkg/middleware"

	"gorm.io/gorm"
)

func SetRouter(ctx context.Context, db *gorm.DB) (*gin.Engine, error) {
	router := gin.Default()
	router.HandleMethodNotAllowed = true

	loanRepo := loanR.NewRepo(db, db)
	repaymentRepo := repaymentR.NewRepo(db, db)
	userRepo := userR.NewRepo(db, db)
	loanSvc := loanS.NewService(loanRepo, repaymentRepo)
	userSvc := userS.NewService(userRepo)
	loanValidator := loanV.NewValidator()
	userValidator := userV.NewValidator()
	loanHandler := handler.NewLoanHandler(loanSvc, loanValidator)
	userHandler := handler.NewUserHandler(userSvc, userValidator)
	authMiddleware := middleware.NewAuthMiddleware(userRepo)
	routerV1 := router.Group("/v1")
	routerV1.POST("/login", userHandler.Login)
	routerV1.POST("/loan", authMiddleware.Authenticate, authMiddleware.Authorize("ADMIN"), loanHandler.CreateLoanHandler)
	routerV1.PUT("/loan/:id", loanHandler.UpdateLoanHandler)
	routerV1.GET("/loan/:id", loanHandler.GetLoanHandler)

	return router, nil
}

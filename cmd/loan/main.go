package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/okaaryanata/loan/internal/api/health"
	"github.com/okaaryanata/loan/internal/api/loan"
	"github.com/okaaryanata/loan/internal/api/repayment"
	"github.com/okaaryanata/loan/internal/api/user"
	"github.com/okaaryanata/loan/internal/app"
	"github.com/okaaryanata/loan/internal/repository"
	"github.com/okaaryanata/loan/internal/service"
)

func main() {
	startService()
}

func startService() {
	app := &app.AppConfig{}
	app.InitService()
	defer app.DB.Close()

	// Repository
	userRepo := repository.NewUserRepository(app.DB)
	loanRepo := repository.NewLoanRepository(app.DB)
	repaymentRepo := repository.NewRepaymentRepository(app.DB)

	// Service
	userSvc := service.NewUserService(userRepo)
	loanSvc := service.NewLoanService(app.DB, loanRepo, repaymentRepo)
	repaymentSvc := service.NewRepaymentService(loanSvc, repaymentRepo)

	// Create controller
	healthController := health.NewHealthController()
	loanController := loan.NewLoanController(loanSvc)
	userController := user.NewUserController(userSvc)
	repaymentController := repayment.NewRepaymentController(repaymentSvc)

	// Create main route
	router := gin.Default()
	mainRoute := router.Group("/svaha-loan")

	// Register routes
	healthController.RegisterRoutes(mainRoute)
	loanController.RegisterRoutes(mainRoute)
	userController.RegisterRoutes(mainRoute)
	repaymentController.RegisterRoutes(mainRoute)

	host := fmt.Sprintf("%s:%s", app.Host, app.Port)
	router.Run(host)
}

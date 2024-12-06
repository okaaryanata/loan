package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/okaaryanata/loan/internal/app"
	"github.com/okaaryanata/loan/internal/helper"
	"github.com/okaaryanata/loan/internal/repository"
	"github.com/okaaryanata/loan/internal/service"
)

func main() {
	fmt.Println("Hello World")

	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init Postgres DB
	dbPool, err := app.InitPostgresDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// repository
	userRepo := repository.NewUserRepository(dbPool)
	loanRepo := repository.NewLoanRepository(dbPool)
	repaymentRepo := repository.NewRepaymentRepository(dbPool)

	// Service
	service.NewUserService(userRepo)
	loanSvc := service.NewLoanService(dbPool, loanRepo, repaymentRepo)
	service.NewRepaymentService(loanSvc, repaymentRepo)

	jakartaTime, err := helper.JakartaTime()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(jakartaTime.AddDate(0, 0, 7))

	defer dbPool.Close()
}

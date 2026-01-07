package main

import (
	"context"
	"log"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/funcBank_Api/config"
	"github.com/funcBank_Api/repository"
	"github.com/funcBank_Api/services"
)

var fundService *services.FundService

func init() {
	config.LoadConfig()
	config.ConnectDb()

	db := config.DB.Client().Database(config.Cfg.DBName)

	fundRepo := repository.NewFundRepo(db)
	fundService = services.NewFundService(fundRepo)
}

func handler(ctx context.Context) error {
	log.Println("Running daily return calculation")

	err := fundService.CalculateReturns()
	if err != nil {
		log.Println("Cron failed:", err)
		return err // enables retry
	}

	log.Println("Cron completed successfully")
	return nil
}

func main() {
	lambda.Start(handler)
}

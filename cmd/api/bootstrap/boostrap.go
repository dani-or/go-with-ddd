package bootstrap

import (
	//"fmt"
	"nequi.com/poc-services/internal/platform/server"
	"nequi.com/poc-services/internal/platform/storage"
	"nequi.com/poc-services/internal/services/credits"
	"nequi.com/poc-services/internal/repository"
)

const (
	host = "0.0.0.0"
	port = 8080
)

func Run() error {
	creditsRepository := dynamo.NewDynamoRepository()
	var _ creditsrepository.CreditsRepository = creditsRepository
	getCreditService := services.NewGetCreditService(creditsRepository)
	srv := server.New(host, port, getCreditService)
	return srv.Run()
}
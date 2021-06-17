package services

import (
	"context"
	"nequi.com/poc-services/internal/repository"
)

type GetCreditsByClientService struct {
	creditsRepository creditsrepository.CreditsRepository
}

func NewGetCreditsByClientService(creditsRepositoryIn creditsrepository.CreditsRepository) GetCreditsByClientService {
	return GetCreditsByClientService{
		creditsRepository: creditsRepositoryIn,
	}
}

func (h GetCreditsByClientService) GetCreditsByClient(ctx context.Context, documentType, documentNumber string) ([]string, error) {
	return h.creditsRepository.GetCreditsByClient(documentType, documentNumber)
}
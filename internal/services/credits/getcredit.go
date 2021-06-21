package services

import (
	"nequi.com/poc-services/internal/repository"
	"nequi.com/poc-services/internal/domain"
)

type GetCreditService struct {
	creditsRepository creditsrepository.CreditsRepository
}

func NewGetCreditService(creditsRepositoryIn creditsrepository.CreditsRepository) GetCreditService {
	return GetCreditService{
		creditsRepository: creditsRepositoryIn,
	}
}

func (h GetCreditService) GetCredit(customerId, debenture string) (credit.Credit, error) {
	//Acá va la logica de mi negocio
	return h.creditsRepository.GetCredit(customerId, debenture)
}
package creditsrepository

import (
	"nequi.com/poc-services/internal/domain"
)

type CreditsRepository interface {
	GetCredit(customerId, debenture string) (credit.Credit, error)
}
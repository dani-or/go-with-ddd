package creditsrepository

import (
)

type CreditsRepository interface {
	GetCreditsByClient(documentType, documentNumber string) ([]string, error)
}
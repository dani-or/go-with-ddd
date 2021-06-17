package dynamo

import (
)

type DynamoRepository struct {
}

func NewDynamoRepository() *DynamoRepository {
	return &DynamoRepository{
	}
}

func (r *DynamoRepository) GetCreditsByClient(documentType, documentNumber string) ([]string, error) {
	return []string{"holaaa"}, nil
}
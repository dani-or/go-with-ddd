package dynamo

import (
	"log"
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
	"nequi.com/poc-services/internal/domain"
	"errors"
)

type DynamoRepository struct {
	client *dynamodb.DynamoDB
}

func NewDynamoRepository() *DynamoRepository {
	svc := dynamodb.New(session.New(),&aws.Config{Region: aws.String("us-east-1")})
	return &DynamoRepository{
		client: svc,
	}
}

//Estructura del item de DB
type Item struct {
    CustomerId   string
    Debenture  string
    EndDate   int
	ProductId string
	Value int
	StartDate int
	Status	int
}

//requiere la variable de entorno export NEQUI_CREDITS_TABLE_NAME=credit-customer-product-qa
func (r *DynamoRepository) GetCredit(customerId, debenture string) (credit.Credit, error) {
	tableName := os.Getenv("NEQUI_CREDITS_TABLE_NAME")
	result, err := r.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"customerId": {
				S: aws.String(customerId),
			},
			"debenture": {
				S: aws.String(debenture),
			},
		},
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
		return credit.Credit{}, err;
	}
	if result.Item == nil {
		fmt.Println( "Could not find '" + customerId + "'")
		return credit.Credit{}, errors.New("No existe ese customerId");
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
		return credit.Credit{}, errors.New("Error de parse");
	}
	
	credit, error := credit.NewCredit(item.Value, item.Status, item.EndDate, item.StartDate , item.Debenture )
	return credit, error
}
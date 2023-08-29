package services

import (
	"fmt"
	"gitlab/assai-clientes/lambda-chapter-products/entities"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type DynamoService struct{}

func (d DynamoService) RegisterLogAudit(logAudit entities.LogAuditEntity) *error {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	client := dynamodb.New(sess)

	item, err := dynamodbattribute.MarshalMap(logAudit)
	if err != nil {
		log.Fatalf("Got error marshalling new log audit item: %s", err)
		return &err
	}

	tableName := "log_audit"

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = client.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
		return &err
	}

	fmt.Println("Successfully added")
	return nil
}

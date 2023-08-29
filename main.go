package main

import (
	"encoding/json"
	"fmt"
	"gitlab/assai-clientes/lambda-chapter-products/entities"
	"gitlab/assai-clientes/lambda-chapter-products/services"
	"gitlab/assai-clientes/lambda-chapter-products/usecases"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Id        string `json:"id"`
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	Target    string `json:"target"`
	Action    string `json:"action"`
	OldValue  string `json:"old_value"`
	NewValue  string `json:"new_value"`
	CreatedAt string `json:"created_at"`
}

func HandleRequest(event events.SQSEvent) {

	dynamoService := services.DynamoService{}

	for _, message := range event.Records {
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)

		var eventBody Event
		if err := json.Unmarshal([]byte(message.Body), &eventBody); err != nil {
			fmt.Println("Erro ao decodificar o corpo da mensagem:", err)
			continue
		}

		logAudit := entities.LogAuditEntity{
			Id:        eventBody.Id,
			UserId:    eventBody.UserId,
			Name:      eventBody.Name,
			Target:    eventBody.Target,
			Action:    eventBody.Action,
			OldValue:  eventBody.OldValue,
			NewValue:  eventBody.NewValue,
			CreatedAt: eventBody.CreatedAt,
		}

		err := usecases.RegisterLogAudit(logAudit, dynamoService)

		if err != nil {
			fmt.Println("lambda executada com sucesso")
		} else {
			fmt.Println("lambda executada com sucesso")
		}
	}
}

func main() {
	lambda.Start(HandleRequest)
}

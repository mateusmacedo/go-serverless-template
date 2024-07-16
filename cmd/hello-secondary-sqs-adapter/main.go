package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"go-sls-template/internal/handlers"
)

func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		log.Printf("Processing message ID: %s", message.MessageId)
		log.Printf("Message body: %s", message.Body)

		helloInput := &handlers.HelloInput{}
		err := json.Unmarshal([]byte(message.Body), helloInput)
		if err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			return err
		}

		helloOutput := handlers.HelloSecondary(*helloInput)
		log.Printf("Hello output: %v", helloOutput)
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}

package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"go-sls-template/internal/handlers"
	"go-sls-template/pkg/aws/sqs"
)

func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		body := &sqs.Body{}
		if err := json.Unmarshal([]byte(message.Body), body); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			return err
		}

		msg := &handlers.HelloInput{}
		if err := json.Unmarshal([]byte(body.Message), msg); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			return err
		}

		output := handlers.HelloSecondary(*msg)

		log.Printf("Hello output: %v", output.Message)
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}

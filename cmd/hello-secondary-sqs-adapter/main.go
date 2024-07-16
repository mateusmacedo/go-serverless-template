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
		if err := processMessage(message); err != nil {
			log.Printf("Failed to process message: %v", err)
			return err
		}
	}
	return nil
}

func processMessage(message events.SQSMessage) error {
	body := &sqs.Body{}
	if err := json.Unmarshal([]byte(message.Body), body); err != nil {
		return logAndReturnError("Failed to unmarshal message body", err)
	}

	msg := &handlers.HelloInput{}
	if err := json.Unmarshal([]byte(body.Message), msg); err != nil {
		return logAndReturnError("Failed to unmarshal inner message", err)
	}

	output := handlers.HelloSecondary(*msg)
	log.Printf("Hello output: %v", output.Message)
	return nil
}

func logAndReturnError(message string, err error) error {
	log.Printf("%s: %v", message, err)
	return err
}

func main() {
	lambda.Start(Handler)
}

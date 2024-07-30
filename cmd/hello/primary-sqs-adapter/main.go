package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		log.Printf("Received message: %s", message.Body)
		// Process the message
		fmt.Println("Message processed:", message.Body)
	}
	return nil
}

func main() {
	lambda.Start(handler)
}

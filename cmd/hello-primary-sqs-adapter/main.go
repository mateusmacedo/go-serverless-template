package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"go-sls-template/internal/hello/application"
	"go-sls-template/internal/hello/domain"
)

type body struct {
	Type             string `json:"Type"`
	MessageId        string `json:"MessageId"`
	TopicArn         string `json:"TopicArn"`
	Message          string `json:"Message"`
	UnsubscribeURL   string `json:"UnsubscribeURL"`
	SignatureVersion string `json:"SignatureVersion"`
	Signature        string `json:"Signature"`
	SigningCertURL   string `json:"SigningCertURL"`
}

type sqsAdapter struct {
	handler *application.HelloHandler
}

func newSqsAdapter(handler *application.HelloHandler) *sqsAdapter {
	return &sqsAdapter{handler: handler}
}

func logAndReturnError(message string, err error) error {
	log.Printf("%s: %v", message, err)
	return err
}

func (h *sqsAdapter) adapt(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, record := range sqsEvent.Records {
		body := &body{}
		if err := json.Unmarshal([]byte(record.Body), body); err != nil {
			return logAndReturnError("Failed to unmarshal message body", err)
		}

		input := application.DispactherInputMsg{}
		if err := json.Unmarshal([]byte(body.Message), &input); err != nil {
			return logAndReturnError("Failed to unmarshal message content", err)
		}

		helloHandlerMsg := application.HandlerInputMsg{
			Name:   input.Message,
			Suffix: "from primary",
		}

		output := h.handler.Handle(ctx, helloHandlerMsg)

		log.Printf("Output: %s", output.Message)
	}

	return nil
}

func main() {
	service := domain.NewHello()
	handler := application.NewHelloHandler(service)
	adapter := newSqsAdapter(handler)

	lambda.Start(adapter.adapt)
}

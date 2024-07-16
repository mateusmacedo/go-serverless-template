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

func (b *body) unmarshal(recordBody string) error {
	return json.Unmarshal([]byte(recordBody), b)
}

func (b *body) ToDispatcherHandlerInputMsg(recordBody string) (application.DispactherHandlerInputMsg, error) {
	err := b.unmarshal(recordBody)
	if err != nil {
		return application.DispactherHandlerInputMsg{}, err
	}
	var input application.DispactherHandlerInputMsg
	err = json.Unmarshal([]byte(b.Message), &input)
	return input, err
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
		var b body
		input, err := b.ToDispatcherHandlerInputMsg(record.Body)
		if err != nil {
			return logAndReturnError("Failed to unmarshal message content", err)
		}

		helloHandlerMsg := application.HelloHandleInputMsg{
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

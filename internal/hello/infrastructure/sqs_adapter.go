package infrastructure

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"

	"go-sls-template/internal/hello/application"
)

func logAndReturnError(message string, err error) error {
	log.Printf("%s: %v", message, err)
	return err
}

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

func (b *body) unmarshal(recordBody []byte) error {
	return json.Unmarshal(recordBody, b)
}

func (b *body) toDispatcherHandlerInputMsg(recordBody []byte) (application.DispactherHandlerInputMsg, error) {
	err := b.unmarshal(recordBody)
	if err != nil {
		return application.DispactherHandlerInputMsg{}, err
	}
	var input application.DispactherHandlerInputMsg
	err = json.Unmarshal([]byte(b.Message), &input)
	return input, err
}

type sqsAdapter struct {
	name    string
	handler *application.HelloHandler
}

func NewSqsAdapter(handler *application.HelloHandler, name string) *sqsAdapter {
	return &sqsAdapter{
		name:    name,
		handler: handler,
	}
}

func (h *sqsAdapter) Adapt(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, record := range sqsEvent.Records {
		if len(record.Body) == 0 {
			return logAndReturnError("Record body is empty", nil)
		}

		var b body
		input, err := b.toDispatcherHandlerInputMsg([]byte(record.Body))
		if err != nil {
			return logAndReturnError("Failed to unmarshal message content: "+record.Body, err)
		}

		helloHandlerMsg := application.HelloHandleInputMsg{
			Name:   input.Message,
			Suffix: h.name,
		}

		output := h.handler.Handle(ctx, helloHandlerMsg)
		log.Printf("Output: %s", output.Message)
	}

	return nil
}

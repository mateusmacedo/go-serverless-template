package infrastructure

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"go.uber.org/zap"

	"go-sls-template/internal/hello/application"
)

var logger, _ = zap.NewProduction()

func logAndReturnError(message string, err error) error {
	logger.Error(message, zap.Error(err))
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
	if err := b.unmarshal(recordBody); err != nil {
		return application.DispactherHandlerInputMsg{}, errors.New("failed to unmarshal body: " + string(recordBody))
	}

	var input application.DispactherHandlerInputMsg
	if err := json.Unmarshal([]byte(b.Message), &input); err != nil {
		return application.DispactherHandlerInputMsg{}, errors.New("failed to unmarshal message content: " + b.Message)
	}

	return input, nil
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
			return logAndReturnError("Failed to process record", err)
		}

		helloHandlerMsg := application.HelloHandleInputMsg{
			Name:   input.Message,
			Suffix: h.name,
		}

		output := h.handler.Handle(ctx, helloHandlerMsg)
		logger.Info("Message processed successfully", zap.String("output", output.Message))
	}

	return nil
}

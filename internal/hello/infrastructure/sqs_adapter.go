package infrastructure

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"

	"go-sls-template/internal/hello/application"
	pkg_app "go-sls-template/pkg/application"
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

func (b *body) unmarshal(recordBody []byte) error {
	return json.Unmarshal(recordBody, b)
}

func (b *body) toDispatcherHandlerInputMsg(recordBody []byte) (application.DispactherInput, error) {
	if err := b.unmarshal(recordBody); err != nil {
		return application.DispactherInput{}, errors.New("failed to unmarshal body: " + string(recordBody))
	}

	var input application.DispactherInput
	if err := json.Unmarshal([]byte(b.Message), &input); err != nil {
		return application.DispactherInput{}, errors.New("failed to unmarshal message content: " + b.Message)
	}

	return input, nil
}

type sqsAdapter struct {
	name    string
	handler *application.HelloHandler
	logger  pkg_app.Logger
}

func NewSqsAdapter(handler *application.HelloHandler, name string, logger pkg_app.Logger) *sqsAdapter {
	return &sqsAdapter{
		name:    name,
		handler: handler,
		logger:  logger,
	}
}

func (h *sqsAdapter) logAndReturnError(logger pkg_app.Logger, message string, err error) error {
	logger.Error(message, err)
	return err
}

func (h *sqsAdapter) Adapt(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, record := range sqsEvent.Records {
		if len(record.Body) == 0 {
			return h.logAndReturnError(h.logger, "Record body is empty", nil)
		}

		var b body
		dispatcherInput, err := b.toDispatcherHandlerInputMsg([]byte(record.Body))
		if err != nil {
			return h.logAndReturnError(h.logger, "Failed to process record", err)
		}

		handlerInput := application.HelloHandleInputMsg{
			Name:   dispatcherInput.Message,
			Suffix: h.name,
			ID:     record.MessageId,
		}

		handlerOutput, err := h.handler.Handle(ctx, handlerInput)
		if err != nil {
			return h.logAndReturnError(h.logger, "Failed to process message", err)
		}

		h.logger.Info("Message processed successfully", "output", handlerOutput.Message)
	}

	return nil
}

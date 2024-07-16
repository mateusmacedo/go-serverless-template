package infrastructure

import (
	"context"

	"github.com/aws/aws-lambda-go/events"

	"go-sls-template/internal/hello/application"
	pkg_app "go-sls-template/pkg/application"
)

func createErrorResponse(logger pkg_app.Logger, message string, statusCode int, err error) (events.APIGatewayProxyResponse, error) {
	logger.Error(message, err)
	return events.APIGatewayProxyResponse{Body: message, StatusCode: statusCode}, err
}

type httpAdapter struct {
	handler *application.DispactherHandler
	logger  pkg_app.Logger
}

func NewHttpAdapter(handler *application.DispactherHandler, logger pkg_app.Logger) *httpAdapter {
	return &httpAdapter{handler: handler, logger: logger}
}

func (h *httpAdapter) Adapt(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]
	if name == "" {
		return createErrorResponse(h.logger, "Name parameter is missing", 400, nil)
	}

	msg := application.DispactherHandlerInputMsg{
		Message: name,
	}

	err := h.handler.Handle(ctx, msg)
	if err != nil {
		return createErrorResponse(h.logger, "Failed to handle message", 500, err)
	}

	h.logger.Info("Message handled successfully", "name", name)
	return events.APIGatewayProxyResponse{Body: "Message handled", StatusCode: 200}, nil
}

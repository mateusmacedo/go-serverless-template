package infrastructure

import (
	"context"

	"github.com/aws/aws-lambda-go/events"

	"go-sls-template/internal/hello/application"
	pkg_app "go-sls-template/pkg/application"
)

type httpAdapter struct {
	handler *application.DispactherHandler
	logger  pkg_app.Logger
}

func NewHttpAdapter(handler *application.DispactherHandler, logger pkg_app.Logger) *httpAdapter {
	return &httpAdapter{handler: handler, logger: logger}
}

func (h *httpAdapter) createErrorResponse(message string, statusCode int, err error) (events.APIGatewayProxyResponse, error) {
	h.logger.Error(message, err)
	return events.APIGatewayProxyResponse{Body: message, StatusCode: statusCode}, err
}

func (h *httpAdapter) createSuccessResponse(body interface{}, statusCode int, message string, keyValues ...interface{}) (events.APIGatewayProxyResponse, error) {
	h.logger.Info(message, keyValues...)
	return events.APIGatewayProxyResponse{Body: body.(string), StatusCode: statusCode}, nil
}

func (h *httpAdapter) Adapt(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]
	if name == "" {
		return h.createErrorResponse("Name parameter is missing", 400, nil)
	}

	input := application.DispactherHandlerInputMsg{
		Message: name,
	}

	err := h.handler.Handle(ctx, input)
	if err != nil {
		return h.createErrorResponse("Failed to handle message", 500, err)
	}

	return h.createSuccessResponse("Hello, "+name, 200, "Successfully handled message", "name", name)
}

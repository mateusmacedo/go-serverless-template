package aws

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"

	"go-sls-template/internal/hello/application"
)

func createErrorResponse(message string, statusCode int, err error) (events.APIGatewayProxyResponse, error) {
	log.Println(message, err)
	return events.APIGatewayProxyResponse{Body: message, StatusCode: statusCode}, err
}

type httpAdapter struct {
	handler *application.DispactherHandler
}

func NewHttpAdapter(handler *application.DispactherHandler) *httpAdapter {
	return &httpAdapter{handler: handler}
}

func (h *httpAdapter) Adapt(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]
	msg := application.DispactherHandlerInputMsg{
		Message: name,
	}

	err := h.handler.Handle(ctx, msg)
	if err != nil {
		return createErrorResponse("Failed to handle message", 500, err)
	}

	return events.APIGatewayProxyResponse{Body: "Message handled", StatusCode: 200}, nil
}

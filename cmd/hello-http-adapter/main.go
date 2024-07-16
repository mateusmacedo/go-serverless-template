package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"

	"go-sls-template/internal/hello/application"
	"go-sls-template/pkg/infrastructure"
)

type httpAdapter struct {
	handler *application.DispactherHandler
}

func newHttpAdapter(handler *application.DispactherHandler) *httpAdapter {
	return &httpAdapter{handler: handler}
}

func createErrorResponse(message string, statusCode int, err error) (events.APIGatewayProxyResponse, error) {
	log.Println(message, err)
	return events.APIGatewayProxyResponse{Body: message, StatusCode: statusCode}, err
}

func (h *httpAdapter) adapt(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	snsClient := sns.NewFromConfig(cfg)

	topicArn := os.Getenv("AWS_SNS_HELLO_TOPIC")
	if topicArn == "" {
		log.Fatalf("AWS_SNS_HELLO_TOPIC environment variable is not set")
	}

	dispatcher := infrastructure.NewSnsDispatcher[application.DispactherHandlerInputMsg](snsClient, topicArn)
	handler := application.NewDispactherHandler(dispatcher)
	adapter := newHttpAdapter(handler)

	lambda.Start(adapter.adapt)
}

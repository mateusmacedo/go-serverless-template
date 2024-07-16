package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"go-sls-template/internal/handlers"
)

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]
	rawmessage := handlers.HelloSecondary(name)
	message := fmt.Sprintf(" { \"Message\" : \"%s\" }", rawmessage)
	return events.APIGatewayProxyResponse{Body: message, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}

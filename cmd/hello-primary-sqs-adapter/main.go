package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"go-sls-template/internal/hello/application"
	"go-sls-template/internal/hello/domain"
	"go-sls-template/internal/hello/infrastructure"
)

func main() {
	service := domain.NewHello()
	handler := application.NewHelloHandler(service)
	adapter := infrastructure.NewSqsAdapter(handler, "from hello-primary-sqs-adapter")
	lambda.Start(adapter.Adapt)
}

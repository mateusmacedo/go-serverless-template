package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"go-sls-template/internal/hello/application"
	"go-sls-template/internal/hello/domain"
	"go-sls-template/internal/hello/infrastructure/aws"
)

func main() {
	service := domain.NewHello()
	handler := application.NewHelloHandler(service)
	adapter := aws.NewSqsAdapter(handler, "from hello-secondary-sqs-adapter")
	lambda.Start(adapter.Adapt)
}

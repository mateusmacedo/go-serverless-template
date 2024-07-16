package main

import (
	"github.com/aws/aws-lambda-go/lambda"

	"go-sls-template/internal/hello/application"
	"go-sls-template/internal/hello/domain"
	"go-sls-template/internal/hello/infrastructure"
	"go-sls-template/pkg/infrastructure/log"
)

func main() {
	logger, _ := log.NewZapLogger()
	service := domain.NewHello()
	handler := application.NewHelloHandler(service)
	adapter := infrastructure.NewSqsAdapter(handler, "from hello-secondary-sqs-adapter", logger)
	lambda.Start(adapter.Adapt)
}

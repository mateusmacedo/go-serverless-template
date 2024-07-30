package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"go-sls-template/internal/hello/application"
	"go-sls-template/internal/hello/domain"
	"go-sls-template/internal/hello/infrastructure"
	"go-sls-template/pkg/infrastructure/log"
)

func main() {
	logger, _ := log.NewZapLogger()
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Error("Error loading AWS config", err)
	}

	dynamoClient := dynamodb.NewFromConfig(cfg)
	tableName := os.Getenv("DYNAMODB_TABLE")
	if tableName == "" {
		logger.Error("DYNAMODB_TABLE environment variable is not set", nil)
	}

	service := domain.NewHello()
	repository := infrastructure.NewDynamoDBRepository(dynamoClient, tableName)
	handler := application.NewHelloHandler(service, repository)
	adapter := infrastructure.NewSqsAdapter(handler, "from hello-primary-sqs-adapter", logger)

	lambda.Start(adapter.Adapt)
}

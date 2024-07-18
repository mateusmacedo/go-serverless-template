package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"

	"go-sls-template/internal/hello/application"
	"go-sls-template/internal/hello/infrastructure"
	"go-sls-template/pkg/infrastructure/aws"
	"go-sls-template/pkg/infrastructure/log"
)

func main() {
	logger, _ := log.NewZapLogger()
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		logger.Error("Error loading AWS config", err)
	}

	snsClient := sns.NewFromConfig(cfg)

	topicArn := os.Getenv("AWS_SNS_HELLO_TOPIC")
	if topicArn == "" {
		logger.Error("AWS_SNS_HELLO_TOPIC environment variable is not set", nil)

	}

	dispatcher := aws.NewSnsDispatcher[application.DispactherHandlerInputMsg](snsClient, topicArn)
	handler := application.NewDispactherHandler(dispatcher)
	adapter := infrastructure.NewHttpAdapter(handler, logger)

	lambda.Start(adapter.Adapt)
}

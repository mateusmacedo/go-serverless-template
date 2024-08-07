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
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Error("Error loading AWS config", err)
	}

	snsClient := sns.NewFromConfig(cfg)

	topicArn := os.Getenv("SNS_TOPIC")
	if topicArn == "" {
		logger.Error("SNS_TOPIC environment variable is not set", nil)

	}

	dispatcher := aws.NewSnsDispatcher[application.DispactherInput](snsClient, topicArn)
	handler := application.NewDispacther(dispatcher)
	adapter := infrastructure.NewHttpAdapter(handler, logger)

	lambda.Start(adapter.Adapt)
}

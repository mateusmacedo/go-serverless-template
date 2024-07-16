package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"

	"go-sls-template/internal/hello/application"
	"go-sls-template/internal/hello/infrastructure"
	"go-sls-template/pkg/infrastructure/aws"
)

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

	dispatcher := aws.NewSnsDispatcher[application.DispactherHandlerInputMsg](snsClient, topicArn)
	handler := application.NewDispactherHandler(dispatcher)
	adapter := infrastructure.NewHttpAdapter(handler)

	lambda.Start(adapter.Adapt)
}

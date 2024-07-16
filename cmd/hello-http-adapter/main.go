package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"

	"go-sls-template/internal/handlers"
)

var (
	snsClient *sns.Client
	topicArn  string
)

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	snsClient = sns.NewFromConfig(cfg)

	topicArn = os.Getenv("AWS_SNS_HELLO_TOPIC")
	if topicArn == "" {
		log.Fatalf("AWS_SNS_HELLO_TOPIC environment variable is not set")
	}
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]

	messageToPublish, err := json.Marshal(handlers.HelloInput{Name: name})
	if err != nil {
		return createErrorResponse("Failed to marshal message", 500, err)
	}

	if err = publishMessage(ctx, string(messageToPublish)); err != nil {
		return createErrorResponse("Failed to publish message", 500, err)
	}

	messageToRespond, err := json.Marshal(handlers.HelloOutput{Message: "Broadcasted hello to " + name})
	if err != nil {
		return createErrorResponse("Failed to marshal response", 500, err)
	}

	return events.APIGatewayProxyResponse{Body: string(messageToRespond), StatusCode: 200}, nil
}

func publishMessage(ctx context.Context, message string) error {
	input := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(topicArn),
	}

	_, err := snsClient.Publish(ctx, input)
	return err
}

func createErrorResponse(message string, statusCode int, err error) (events.APIGatewayProxyResponse, error) {
	log.Println(message, err)
	return events.APIGatewayProxyResponse{Body: message, StatusCode: statusCode}, err
}

func main() {
	lambda.Start(Handler)
}

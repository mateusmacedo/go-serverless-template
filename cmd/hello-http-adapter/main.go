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

var snsClient *sns.Client
var topicArn string

func init() {
	// Carregar configuração do SDK da AWS
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Criar um cliente SNS
	snsClient = sns.NewFromConfig(cfg)

	// Carregar ARN do tópico SNS da variável de ambiente
	topicArn = os.Getenv("AWS_SNS_HELLO_TOPIC")
	if topicArn == "" {
		log.Fatalf("AWS_SNS_HELLO_TOPIC environment variable is not set")
	}
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]

	messageToPublish, err := json.Marshal(handlers.HelloInput{Name: name})
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Failed to marshal message", StatusCode: 500}, err
	}

	err = publishMessage(ctx, string(messageToPublish))
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Failed to publish message", StatusCode: 500}, err
	}

	messageToRespond, err := json.Marshal(handlers.HelloOutput{Message: "Broadcasted hello to " + name})
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Failed to marshal response", StatusCode: 500}, err
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

func main() {
	lambda.Start(Handler)
}

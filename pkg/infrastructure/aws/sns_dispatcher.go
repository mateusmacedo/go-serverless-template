package aws

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"

	"go-sls-template/pkg/application"
)

type SnsDispatcher[T any] struct {
	client   *sns.Client
	topicArn string
}

func NewSnsDispatcher[T any](client *sns.Client, topicArn string) application.MessageDispatcher[T] {
	return &SnsDispatcher[T]{
		client:   client,
		topicArn: topicArn,
	}
}

func (p *SnsDispatcher[T]) Dispatch(ctx context.Context, message T) error {
	messageToPublish, err := json.Marshal(message)
	if err != nil {
		return err
	}

	input := &sns.PublishInput{
		Message:  aws.String(string(messageToPublish)),
		TopicArn: aws.String(p.topicArn),
	}

	_, err = p.client.Publish(ctx, input)
	return err
}

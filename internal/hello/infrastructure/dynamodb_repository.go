package infrastructure

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	"go-sls-template/internal/hello/domain"
)

type dynamoDBRepository struct {
	db        *dynamodb.Client
	tableName string
}

func NewDynamoDBRepository(db *dynamodb.Client, tableName string) domain.MessageRepository {
	return &dynamoDBRepository{
		db:        db,
		tableName: tableName,
	}
}

func (r *dynamoDBRepository) SaveMessage(ctx context.Context, message domain.Message) error {
	av, err := attributevalue.MarshalMap(message)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      av,
	}

	_, err = r.db.PutItem(ctx, input)
	return err
}

func (r *dynamoDBRepository) GetMessage(ctx context.Context, id string) (*domain.Message, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	}

	result, err := r.db.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	var message domain.Message
	err = attributevalue.UnmarshalMap(result.Item, &message)
	return &message, err
}

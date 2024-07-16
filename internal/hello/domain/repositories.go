package domain

import "context"

type MessageRepository interface {
	SaveMessage(ctx context.Context, message Message) error
	GetMessage(ctx context.Context, id string) (*Message, error)
}

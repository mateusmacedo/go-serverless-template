package application

import "context"

type MessageDispatcher[T any] interface {
	Dispatch(ctx context.Context, message T) error
}

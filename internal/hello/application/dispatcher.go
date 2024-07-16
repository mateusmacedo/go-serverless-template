package application

import (
	"context"

	"go-sls-template/pkg/application"
)

type DispactherHandlerInputMsg struct {
	Message string `json:"message"`
}

type DispactherHandler struct {
	dispatcher application.MessageDispatcher[DispactherHandlerInputMsg]
}

func NewDispactherHandler(publisher application.MessageDispatcher[DispactherHandlerInputMsg]) *DispactherHandler {
	return &DispactherHandler{
		dispatcher: publisher,
	}
}

func (h DispactherHandler) Handle(ctx context.Context, msg DispactherHandlerInputMsg) error {
	if err := h.dispatcher.Dispatch(ctx, msg); err != nil {
		return err
	}

	return nil
}

package application

import (
	"context"

	"go-sls-template/pkg/application"
)

type DispactherInputMsg struct {
	Message string `json:"message"`
}

type DispactherHandler struct {
	dispatcher application.MessageDispatcher[DispactherInputMsg]
}

func NewDispactherHandler(publisher application.MessageDispatcher[DispactherInputMsg]) *DispactherHandler {
	return &DispactherHandler{
		dispatcher: publisher,
	}
}

func (h DispactherHandler) Handle(ctx context.Context, msg DispactherInputMsg) error {
	if err := h.dispatcher.Dispatch(ctx, msg); err != nil {
		return err
	}

	return nil
}

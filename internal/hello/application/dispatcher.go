package application

import (
	"context"
	"errors"

	"go-sls-template/pkg/application"
)

type DispactherInput struct {
	Message string `json:"message"`
}

type Dispacther struct {
	dispatcher application.MessageDispatcher[DispactherInput]
}

func NewDispacther(publisher application.MessageDispatcher[DispactherInput]) *Dispacther {
	return &Dispacther{
		dispatcher: publisher,
	}
}

func (h Dispacther) Dispatch(ctx context.Context, msg DispactherInput) error {
	if msg.Message == "" {
		return errors.New("message cannot be empty")
	}
	if err := h.dispatcher.Dispatch(ctx, msg); err != nil {
		return err
	}
	return nil
}

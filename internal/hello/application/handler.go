package application

import (
	"context"
	"time"

	"go-sls-template/internal/hello/domain"
)

type HelloHandleInputMsg struct {
	Name   string `json:"name"`
	Suffix string `json:"suffix"`
	ID     string `json:"id"`
}

type HelloHandleOutputMsg struct {
	Message string `json:"message"`
}

type HelloHandler struct {
	hello      domain.Hello
	repository domain.MessageRepository
}

func NewHelloHandler(hello domain.Hello, repo domain.MessageRepository) *HelloHandler {
	return &HelloHandler{
		hello:      hello,
		repository: repo,
	}
}

func (h *HelloHandler) Handle(ctx context.Context, msg HelloHandleInputMsg) (HelloHandleOutputMsg, error) {
	input := domain.HelloInput(struct {
		Name   string `json:"name"`
		Suffix string `json:"suffix"`
	}{
		Name:   msg.Name,
		Suffix: msg.Suffix,
	})

	output, _ := h.hello.Say(input)
	message := domain.Message{
		ID:        msg.ID,
		Content:   output.Message,
		Timestamp: time.Now(),
	}

	err := h.repository.SaveMessage(ctx, message)
	if err != nil {
		return HelloHandleOutputMsg{}, err
	}
	return HelloHandleOutputMsg(output), nil

}

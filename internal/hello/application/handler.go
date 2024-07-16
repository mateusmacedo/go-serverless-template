package application

import (
	"context"

	"go-sls-template/internal/hello/domain"
)

type HelloHandleInputMsg struct {
	Name   string `json:"name"`
	Suffix string `json:"suffix"`
}

type HelloHandleOutputMsg struct {
	Message string `json:"message"`
}

type HelloHandler struct {
	hello domain.Hello
}

func NewHelloHandler(hello domain.Hello) *HelloHandler {
	return &HelloHandler{
		hello: hello,
	}
}

func (h *HelloHandler) Handle(ctx context.Context, msg HelloHandleInputMsg) (HelloHandleOutputMsg, error) {
	input := domain.HelloInput(msg)

	output, err := h.hello.Say(input)
	if err != nil {
		return HelloHandleOutputMsg{}, err
	}
	return HelloHandleOutputMsg(output), nil

}

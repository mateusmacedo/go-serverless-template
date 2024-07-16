package application

import (
	"context"

	"go-sls-template/internal/hello/domain"
)

type HandlerInputMsg struct {
	Name   string `json:"name"`
	Suffix string `json:"suffix"`
}

type HandlerOutputMsg struct {
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

func (h *HelloHandler) Handle(ctx context.Context, msg HandlerInputMsg) HandlerOutputMsg {
	input := domain.HelloInput(msg)

	output := h.hello.Say(input)

	return HandlerOutputMsg(output)
}

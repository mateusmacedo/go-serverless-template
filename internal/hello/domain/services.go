package domain

type Hello struct{}

func NewHello() Hello {
	return Hello{}
}

func (h *Hello) Say(input HelloInput) (HelloOutput, error) {
	message := h.constructMessage(input)
	return HelloOutput{Message: message}, nil
}

func (h *Hello) constructMessage(input HelloInput) string {
	if input.Suffix == "" {
		return "Hello " + input.Name
	}
	return "Hello " + input.Name + " " + input.Suffix
}

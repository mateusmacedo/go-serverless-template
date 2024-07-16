package domain

type Hello struct{}

func NewHello() Hello {
	return Hello{}
}

func (h *Hello) Say(input HelloInput) (HelloOutput, error) {
	if input.Suffix == "" {
		return HelloOutput{
			Message: "Hello " + input.Name,
		}, nil
	}

	return HelloOutput{
		Message: "Hello " + input.Name + " " + input.Suffix,
	}, nil
}

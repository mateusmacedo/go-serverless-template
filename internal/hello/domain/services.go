package domain

type Hello struct{}

func NewHello() Hello {
	return Hello{}
}

func (h *Hello) Say(input HelloInput) HelloOutput {
	if input.Suffix == "" {
		return HelloOutput{
			Message: "Hello " + input.Name,
		}
	}

	return HelloOutput{
		Message: "Hello " + input.Name + " from " + input.Suffix,
	}
}

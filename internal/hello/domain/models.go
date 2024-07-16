package domain

type HelloInput struct {
	Name   string `json:"name"`
	Suffix string `json:"suffix"`
}

type HelloOutput struct {
	Message string `json:"message"`
}

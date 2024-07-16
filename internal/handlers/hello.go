package handlers

type HelloInput struct {
	Name string `json:"name"`
}

type HelloOutput struct {
	Message string `json:"message"`
}

func Hello(HelloRequest HelloInput) HelloOutput {
	return HelloOutput{
		Message: "Hello " + HelloRequest.Name,
	}
}

func HelloSecondary(HelloRequest HelloInput) HelloOutput {
	return HelloOutput{
		Message: "Hello " + HelloRequest.Name + " from secondary",
	}
}

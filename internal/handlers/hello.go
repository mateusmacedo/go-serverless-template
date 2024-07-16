package handlers

type HelloInput struct {
	Name string `json:"name"`
}

type HelloOutput struct {
	Message string `json:"message"`
}

func HelloPrimary(HelloRequest HelloInput) HelloOutput {
	return HelloOutput{
		Message: "Hello " + HelloRequest.Name + " from primary",
	}
}

func HelloSecondary(HelloRequest HelloInput) HelloOutput {
	return HelloOutput{
		Message: "Hello " + HelloRequest.Name + " from secondary",
	}
}

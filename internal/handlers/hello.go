package handlers

type HelloInput struct {
	Name string `json:"name"`
}

type HelloOutput struct {
	Message string `json:"message"`
}

func createHelloMessage(name, suffix string) HelloOutput {
	return HelloOutput{
		Message: "Hello " + name + " from " + suffix,
	}
}

func HelloPrimary(HelloRequest HelloInput) HelloOutput {
	return createHelloMessage(HelloRequest.Name, "primary")
}

func HelloSecondary(HelloRequest HelloInput) HelloOutput {
	return createHelloMessage(HelloRequest.Name, "secondary")
}

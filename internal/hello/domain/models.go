package domain

import "time"

type HelloInput struct {
	Name   string `json:"name"`
	Suffix string `json:"suffix"`
}

type HelloOutput struct {
	Message string `json:"message"`
}

type Message struct {
	ID        string    `dynamodbav:"id"`
	Content   string    `dynamodbav:"content"`
	Timestamp time.Time `dynamodbav:"timestamp"`
}

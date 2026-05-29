package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	username string `json:"username"`
}

func HandleRequest(event MyEvent) (string, error) {
	if event.username == "" {
		return "", fmt.Errorf("Username cannot be empty")
	}
	return fmt.Sprintf("Succesfully called by - %s", event.username), nil
}

func main() {
	lambda.Start(HandleRequest)
}

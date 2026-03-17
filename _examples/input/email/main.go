package main

import (
	"fmt"

	"github.com/indaco/prompti/input"
)

func main() {
	questionPrompt := &input.Config{
		Message:      "What's your email address?",
		Placeholder:  "Please, provide an email address",
		ErrorMsg:     "Email is mandatory",
		ValidateFunc: input.ValidateEmail,
	}

	result, _ := input.Run(questionPrompt)
	fmt.Println(result)
}

package main

import (
	"fmt"

	"github.com/indaco/prompti/input"
)

func main() {
	questionPrompt := &input.Config{
		Message:      "What's your lucky number?",
		Placeholder:  "Please, tell me your lucky number",
		Initial:      "23",
		ErrorMsg:     "Cannot be blank",
		ValidateFunc: input.ValidateInteger,
	}

	result, _ := input.Run(questionPrompt)
	fmt.Println(result)
}

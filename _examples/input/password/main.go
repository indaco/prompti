package main

import (
	"fmt"

	"github.com/indaco/prompti/input"
)

func main() {
	passwordPrompt := &input.Config{
		Message:     "What's  your password?",
		Placeholder: "Please, provide your password",
		ErrorMsg:    "Password is mandatory",
		Password:    true,
	}

	result, _ := input.Run(passwordPrompt)
	fmt.Println(result)
}

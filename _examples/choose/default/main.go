package main

import (
	"fmt"

	"github.com/indaco/prompti/choose"
)

func main() {
	foodSelectionPrompt := &choose.Config{
		Title:    "What do you wanna eat tonight?",
		ErrorMsg: "Please, select your meal.",
	}

	entries := []choose.Item{
		{Name: "pizza", Desc: "It's always pizza time!"},
		{Name: "kebab", Desc: "I feel turkish today, kebab!"},
		{Name: "carbonara", Desc: "Carbonara, NO cream, please!"},
	}

	result, _ := choose.Run(foodSelectionPrompt, entries)
	fmt.Println(result)
}

package main

import (
	"fmt"

	"github.com/indaco/prompti/toggle"
)

func main() {
	result, _ := toggle.Run(&toggle.Config{Question: "Continue?"})
	fmt.Println(result)
}

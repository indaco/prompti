package main

import (
	"fmt"

	"github.com/indaco/prompti/confirm"
)

func main() {
	result, _ := confirm.Run(&confirm.Config{Question: "Continue?"})
	fmt.Println(result)
}

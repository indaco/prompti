// Collapsible detail/summary prompt with custom expand/collapse indicators.
//
// Run from the repository root:
//
//	go run ./examples/04-detail/custom-indicators
package main

import (
	"fmt"

	"github.com/indaco/prompti/detail"
)

var detailConfig = &detail.Config{
	Summary:            "Configuration Options",
	Content:            "verbose: true\nlog_level: debug\nmax_retries: 3",
	CollapsedIndicator: "[+]",
	ExpandedIndicator:  "[-]",
}

func main() {
	expanded, _ := detail.Run(detailConfig)
	fmt.Println("expanded:", expanded)
}

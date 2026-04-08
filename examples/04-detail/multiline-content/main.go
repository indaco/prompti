// Collapsible detail/summary prompt with multiline content.
//
// Run from the repository root:
//
//	go run ./examples/04-detail/multiline-content
package main

import (
	"fmt"

	"github.com/indaco/prompti/detail"
)

var detailConfig = &detail.Config{
	Summary: "Error Details",
	Content: `Status:  503 Service Unavailable
Endpoint: /api/v1/users
Trace ID: abc-123-def-456

The upstream service did not respond within
the configured timeout of 30 seconds.

Suggested actions:
  1. Check service health dashboard
  2. Verify network connectivity
  3. Review recent deployments`,
}

func main() {
	expanded, _ := detail.Run(detailConfig)
	fmt.Println("expanded:", expanded)
}

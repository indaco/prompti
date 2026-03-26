<h1 align="center">
    <span style="font-family: monospace;">prompti</span>
</h1>
<h2 align="center">
Interactive TUI prompts for Go CLI applications, powered by Charm.
</h2>
<p align="center">
  <a href="https://github.com/indaco/prompti/actions/workflows/ci.yml" target="_blank">
    <img src="https://github.com/indaco/prompti/actions/workflows/ci.yml/badge.svg" alt="CI" />
  </a>
  <a href="https://codecov.io/gh/indaco/prompti" target="_blank">
    <img src="https://codecov.io/gh/indaco/prompti/branch/main/graph/badge.svg" alt="Code coverage" />
  </a>
  <a href="https://goreportcard.com/report/github.com/indaco/prompti" target="_blank">
    <img src="https://goreportcard.com/badge/github.com/indaco/prompti" alt="Go Report Card" />
  </a>
  <a href="https://github.com/indaco/prompti/actions/workflows/security.yml" target="_blank">
    <img src="https://github.com/indaco/prompti/actions/workflows/security.yml/badge.svg" alt="Security Scan" />
  </a>
  <a href="https://github.com/indaco/prompti/releases" target="_blank">
    <img src="https://img.shields.io/github/v/tag/indaco/prompti?label=version&sort=semver&color=4c1" alt="version">
  </a>
  <a href="https://pkg.go.dev/github.com/indaco/prompti" target="_blank">
    <img src="https://pkg.go.dev/badge/github.com/indaco/prompti.svg" alt="Go Reference" />
  </a>
  <a href="LICENSE" target="_blank">
    <img src="https://img.shields.io/badge/license-mit-blue?style=flat-square" alt="License" />
  </a>
  <a href="https://www.jetify.com/devbox" target="_blank">
    <img src="https://www.jetify.com/img/devbox/shield_moon.svg" alt="Built with Devbox" />
  </a>
</p>

> **Note**
> This project was previously hosted at `github.com/sveltinio/prompti`. The repository has been transferred to `github.com/indaco/prompti` under the same maintainer.

**prompti** is a collection of interactive TUI prompt components for Go CLI applications, built on the [Charm](https://charm.sh) ecosystem (bubbletea, bubbles, lipgloss).

## Features

- **6 prompt types**: input, choose, confirm, toggle, detail, progressbar
- **Customizable styles**: every component accepts a `Styles` struct for full visual control
- **Sentinel errors**: distinguish user cancellation (`ErrCancelled`) from empty input (`ErrEmpty`)
- **Built-in validation**: ready-to-use validators for alphanumeric, digits, integers, floats, email, and URL

## Install

Requires **Go 1.25** or higher.

```bash
go get github.com/indaco/prompti@latest
```

## Prompts

### Input

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/input-default.gif" alt="Input example">
</details>

`input` is a text input prompt supporting default values, validation (`type ValidateFunc func(string) error`), password echo mode, and customizable styles.

Built-in validation rules:

- alphanumeric
- digits only
- integers
- floats
- email address
- URL

#### Default

```go
package main

import (
	"fmt"

	"github.com/indaco/prompti/input"
)

func main() {
	questionPrompt := &input.Config{
		Message:     "What's the name of your project?",
		Placeholder: "Please, provide a name for your project",
		ErrorMsg:    "Project name is mandatory",
	}

	result, _ := input.Run(questionPrompt)
	fmt.Println(result)
}
```

#### Default value and validation

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/input-initial-value.gif" alt="Input initial value example">
</details>

```go
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
```

#### Custom styles

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/input-styled.gif" alt="Input custom styles example">
</details>

```go
package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"github.com/indaco/prompti/input"
)

func main() {
	questionPrompt := &input.Config{
		Message:     "What's the name of your project?",
		Placeholder: "Please, provide a name for your project",
		ErrorMsg:    "Project name is mandatory",
		Styles: input.Styles{
			PrefixIcon:      "*",
			PrefixIconColor: compat.AdaptiveColor{Light: lipgloss.Color("#ef4444"), Dark: lipgloss.Color("#ef4444")},
			PlaceholderStyle: lipgloss.NewStyle().
				Background(compat.AdaptiveColor{Light: lipgloss.Color("#3b82f6"), Dark: lipgloss.Color("#3b82f6")}).
				Foreground(compat.AdaptiveColor{Light: lipgloss.Color("#fde68a"), Dark: lipgloss.Color("#fffbeb")}),
		},
	}

	result, _ := input.Run(questionPrompt)
	fmt.Println(result)
}
```

#### Email

```go
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
```

#### Password

```go
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
```

### Choose

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/choose-default.gif" alt="Choose example">
</details>

`choose` is a single-select list prompt for browsing a set of items.

#### Default

```go
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
```

#### Custom styles

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/choose-styled.gif" alt="Choose custom styles example">
</details>

```go
package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"github.com/indaco/prompti/choose"
)

var (
	amber  = compat.AdaptiveColor{Light: lipgloss.Color("#f59e0b"), Dark: lipgloss.Color("#fbbf24")}
	purple = compat.AdaptiveColor{Light: lipgloss.Color("#7e22ce"), Dark: lipgloss.Color("#a855f7")}
	green  = compat.AdaptiveColor{Light: lipgloss.Color("#166534"), Dark: lipgloss.Color("#22c55e")}

	myCustomStyle = choose.Styles{
		PrefixIcon:        "★",
		TitleStyle:        lipgloss.NewStyle().Background(green).Foreground(purple).Padding(0, 1),
		TitleBarStyle:     lipgloss.NewStyle(),
		ItemIcon:          "#",
		ItemStyle:         lipgloss.NewStyle().Foreground(amber),
		SelectedItemStyle: lipgloss.NewStyle().Foreground(purple),
	}
)

func main() {
	foodSelectionPrompt := &choose.Config{
		Title:    "What do you wanna eat tonight?",
		ErrorMsg: "Please, select your meal.",
		ShowHelp: true,
		Styles:   myCustomStyle,
	}

	entries := []choose.Item{
		{Name: "pizza", Desc: "It's always pizza time!"},
		{Name: "kebab", Desc: "I feel turkish today, kebab!"},
		{Name: "carbonara", Desc: "Carbonara, NO cream, please!"},
	}

	result, _ := choose.Run(foodSelectionPrompt, entries)
	fmt.Println(result)
}
```

### Confirm

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/confirm-default.gif" alt="Confirm example">
</details>

`confirm` is a yes/no confirmation dialog rendered in a styled box.

#### Default

```go
package main

import (
	"fmt"

	"github.com/indaco/prompti/confirm"
)

func main() {
	result, _ := confirm.Run(&confirm.Config{Question: "Continue?"})
	fmt.Println(result)
}
```

#### Custom styles

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/confirm-styled.gif" alt="Confirm custom styles example">
</details>

```go
package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"github.com/indaco/prompti/confirm"
)

var (
	cyan  = compat.AdaptiveColor{Light: lipgloss.Color("#4f46e5"), Dark: lipgloss.Color("#c7d2fe")}
	green = compat.AdaptiveColor{Light: lipgloss.Color("#166534"), Dark: lipgloss.Color("#22c55e")}

	infoText = `Lorem ipsum dolor sit amet,
consectetur adipiscing elit %s...`

	Green   = lipgloss.NewStyle().Foreground(green).Render
	message = fmt.Sprintf(infoText, Green("elit"))

	myCustomStyle = confirm.Styles{
		Width:       60,
		BorderColor: cyan,
	}

	confirmConfig = &confirm.Config{
		Message:  message,
		Question: "Continue?",
		Styles:   myCustomStyle,
	}
)

func main() {
	result, _ := confirm.Run(confirmConfig)
	fmt.Println(result)
}
```

### Toggle

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/toggle-default.gif" alt="Toggle example">
</details>

`toggle` is an inline yes/no prompt. It works like `confirm` but renders the options inline rather than in a box.

#### Default

```go
package main

import (
	"fmt"

	"github.com/indaco/prompti/toggle"
)

func main() {
	result, _ := toggle.Run(&toggle.Config{Question: "Continue?"})
	fmt.Println(result)
}
```

#### Custom styles

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/toggle-styled.gif" alt="Toggle custom styles example">
</details>

```go
package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"github.com/indaco/prompti/toggle"
)

var (
	cyan   = compat.AdaptiveColor{Light: lipgloss.Color("#4f46e5"), Dark: lipgloss.Color("#c7d2fe")}
	green  = compat.AdaptiveColor{Light: lipgloss.Color("#166534"), Dark: lipgloss.Color("#22c55e")}
	red    = compat.AdaptiveColor{Light: lipgloss.Color("#ef4444"), Dark: lipgloss.Color("#ef4444")}
	purple = compat.AdaptiveColor{Light: lipgloss.Color("#7e22ce"), Dark: lipgloss.Color("#a855f7")}

	myCustomStyle = toggle.Styles{
		PrefixIcon:        "★",
		PrefixIconColor:   red,
		DialogStyle:       lipgloss.NewStyle().Margin(1, 0),
		ButtonStyle:       lipgloss.NewStyle().Bold(true).Foreground(cyan),
		ActiveButtonStyle: lipgloss.NewStyle().Foreground(green),
	}

	toggleConfig = &toggle.Config{
		Question:          "How do you feel?",
		OkButtonLabel:     "I'm super ok",
		CancelButtonLabel: "Next question, please!",
		Divider:           "|",
		Styles:            myCustomStyle,
	}
)

func main() {
	result, _ := toggle.Run(toggleConfig)
	fmt.Println(result)
}
```

### Detail

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/detail-default.gif" alt="Detail example">
</details>

`detail` is a collapsible detail/summary section, similar to the HTML `<details>` element. It displays a summary line with an expand/collapse indicator and toggles content visibility on user interaction.

#### Default

```go
package main

import (
	"fmt"

	"github.com/indaco/prompti/detail"
)

func main() {
	expanded, err := detail.Run(&detail.Config{
		Summary: "What is prompti?",
		Content: "prompti is a collection of interactive terminal UI prompts\nbuilt on the charmbracelet bubbletea framework.",
	})
	fmt.Println("expanded:", expanded, "err:", err)
}
```

#### Custom indicators

```go
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
```

#### Custom styles

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/detail-styled.gif" alt="Detail custom styles example">
</details>

```go
package main

import (
	"fmt"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/compat"
	"github.com/indaco/prompti/detail"
)

var (
	cyan   = compat.AdaptiveColor{Light: lipgloss.Color("#4f46e5"), Dark: lipgloss.Color("#c7d2fe")}
	green  = compat.AdaptiveColor{Light: lipgloss.Color("#166534"), Dark: lipgloss.Color("#22c55e")}
	red    = compat.AdaptiveColor{Light: lipgloss.Color("#ef4444"), Dark: lipgloss.Color("#ef4444")}
	purple = compat.AdaptiveColor{Light: lipgloss.Color("#7e22ce"), Dark: lipgloss.Color("#a855f7")}

	myCustomStyle = detail.Styles{
		PrefixIcon:      "★",
		PrefixIconColor: red,
		SummaryStyle:    lipgloss.NewStyle().Bold(true).Foreground(purple).PaddingRight(1),
		IndicatorStyle:  lipgloss.NewStyle().Foreground(green).PaddingRight(1),
		ContentStyle:    lipgloss.NewStyle().Foreground(cyan).PaddingLeft(3).MarginTop(1),
		DialogStyle:     lipgloss.NewStyle().Margin(1, 0),
	}

	detailConfig = &detail.Config{
		Summary: "Release Notes v0.3.0",
		Content: "- Added detail/summary component\n- Improved theme system\n- Bug fixes and performance improvements",
		Styles:  myCustomStyle,
	}
)

func main() {
	expanded, _ := detail.Run(detailConfig)
	fmt.Println("expanded:", expanded)
}
```

#### Multiline content

```go
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
```

### ProgressBar

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/progressbar-default.gif" alt="ProgressBar example">
</details>

`progressbar` is an animated progress meter that iterates over a list of items.

#### Default

```go
package main

import (
	"fmt"
	"os"

	"github.com/indaco/prompti/progressbar"
)

func main() {
	fruits := []string{
		"apple",
		"banana",
		"orange",
		"grapes",
		"mellon",
		"strawberry",
		"mango",
		"lemon",
		"apricot",
		"peach",
		"papaya",
		"kiwi",
		"pear",
		"guava",
		"almond",
		"coconut",
		"blackberry",
		"cherry",
		"grapes",
	}

	pbConfig := &progressbar.Config{Items: fruits}

	if err := progressbar.Run(pbConfig); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
```

#### Custom styles

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/progressbar-styled.gif" alt="ProgressBar custom styles example">
</details>

```go
package main

import (
	"fmt"
	"os"

	"charm.land/lipgloss/v2"
	"github.com/indaco/prompti/progressbar"
)

func main() {
	fruits := []string{
		"apple",
		"banana",
		"orange",
		"grapes",
		"mellon",
		"strawberry",
		"mango",
		"lemon",
		"apricot",
		"peach",
		"papaya",
		"kiwi",
		"pear",
		"guava",
		"almond",
		"coconut",
		"blackberry",
		"cherry",
		"grapes",
	}

	pbConfig := &progressbar.Config{
		Items:         fruits,
		OnProgressMsg: "Eating:",
		Styles: progressbar.Styles{
			CurrentItemStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("411")),
			ShowLabel:        true,
			GradientFrom:     lipgloss.Color("#FF7CCB"),
			GradientTo:       lipgloss.Color("#FDFF8C"),
		}}

	if err := progressbar.Run(pbConfig); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
```

#### Concurrent execution

```go
package main

import (
	"fmt"
	"os"

	"github.com/indaco/prompti/progressbar"
)

func main() {
	fruits := []string{
		"apple",
		"banana",
		"orange",
		"grapes",
		"mellon",
		"strawberry",
		"mango",
		"lemon",
		"apricot",
		"peach",
		"papaya",
		"kiwi",
		"pear",
		"guava",
		"almond",
		"coconut",
		"blackberry",
		"cherry",
		"grapes",
	}

	pbConfig := &progressbar.Config{Items: fruits, RunConcurrently: true}

	fmt.Println("Run commands concurrently")
	if err := progressbar.Run(pbConfig); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

---

Made with [Charm](https://charm.sh).

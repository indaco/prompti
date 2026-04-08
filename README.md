<h1 align="center">
    <span style="font-family: monospace;">prompti</span>
</h1>
<h2 align="center">
Build polished, interactive CLI prompts in Go - with minimal setup.
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

<p align="center">
  <b><a href="#quick-start">Quick Start</a></b> |
  <b><a href="#features">Features</a></b> |
  <b><a href="#components">Components</a></b> |
  <b><a href="#usage-modes">Usage Modes</a></b> |
  <b><a href="#examples">Examples</a></b> |
  <b><a href="#related">Related</a></b>
</p>

> **Note**
> This project was previously hosted at `github.com/sveltinio/prompti`. The repository has been transferred to `github.com/indaco/prompti` under the same maintainer.

**prompti** is a collection of interactive TUI prompt components for Go CLI applications, built on the [Charm](https://charm.sh) ecosystem (bubbletea, bubbles, lipgloss).

## Quick Start

Requires **Go 1.25** or higher.

```bash
go get github.com/indaco/prompti@latest
```

```go
package main

import (
	"fmt"
	"log"

	"github.com/indaco/prompti/input"
)

func main() {
	result, err := input.Run(&input.Config{
		Message:     "What's the name of your project?",
		Placeholder: "my-awesome-project",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Project:", result)
}
```

## Features

**Components** - Input, Choose, Confirm (dialog + inline), Detail, ProgressBar

**Styling** - Full [Lipgloss](https://github.com/charmbracelet/lipgloss) v2 style overrides with adaptive light/dark theme

**Validation** - Built-in rules for alphanumeric, digits, integers, floats, email, URL - plus custom `ValidateFunc`

**Integration** - [`huh.Field`](https://github.com/charmbracelet/huh) adapters for multi-step forms

**Accessibility** - Every `huh.Field` adapter implements `RunAccessible` for screen reader support

**Error Handling** - Typed sentinel errors (`ErrCancelled`, `ErrEmpty`) for clean control flow

## Components

| Component | Description | Returns | `huh.Field` |
|-----------|-------------|---------|:-----------:|
| [`input`](#input) | Single-line text entry with optional validation | `string, error` | Yes |
| [`choose`](#choose) | Single-select list from a set of items | `string, error` | Yes |
| [`confirm`](#confirm) | Yes/no prompt - dialog box or inline toggle | `bool, error` | Yes |
| [`detail`](#detail) | Collapsible detail/summary section | `bool, error` | Yes |
| [`progressbar`](#progressbar) | Animated progress meter over a list of items | `error` | No |

---

### Input

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/input-default.gif" alt="Input example">
</details>

Single-line text entry. Returns the entered string, or `ErrCancelled` / `ErrEmpty`.

| Field | Type | Description |
|-------|------|-------------|
| `Message` | `string` | Prompt label shown above the input |
| `Placeholder` | `string` | Ghost text when empty |
| `Initial` | `string` | Pre-filled default value |
| `ErrorMsg` | `string` | Message shown on validation failure |
| `Password` | `bool` | Mask input characters |
| `ValidateFunc` | `ValidateFunc` | Custom validation function |
| `Styles` | `Styles` | Override default appearance |

#### Default

```go
result, err := input.Run(&input.Config{
    Message:     "What's the name of your project?",
    Placeholder: "Please, provide a name for your project",
    ErrorMsg:    "Project name is mandatory",
})
```

#### With default value and validation

```go
result, err := input.Run(&input.Config{
    Message:      "What's your lucky number?",
    Placeholder:  "Please, tell me your lucky number",
    Initial:      "23",
    ErrorMsg:     "Cannot be blank",
    ValidateFunc: input.ValidateInteger,
})
```

#### Email validation

```go
result, err := input.Run(&input.Config{
    Message:      "What's your email address?",
    Placeholder:  "Please, provide an email address",
    ErrorMsg:     "Email is mandatory",
    ValidateFunc: input.ValidateEmail,
})
```

#### Password mode

```go
result, err := input.Run(&input.Config{
    Message:     "What's your password?",
    Placeholder: "Please, provide your password",
    ErrorMsg:    "Password is mandatory",
    Password:    true,
})
```

<details>
<summary><b>Custom styles</b></summary>

```go
result, err := input.Run(&input.Config{
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
})
```

</details>

---

### Choose

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/choose-default.gif" alt="Choose example">
</details>

Single-select list for browsing and picking from a set of items.

| Field | Type | Description |
|-------|------|-------------|
| `Title` | `string` | Prompt label above the list |
| `ListHeight` | `int` | Max visible items (auto-calculated if 0) |
| `DefaultWidth` | `int` | List width |
| `ErrorMsg` | `string` | Message shown when no selection is made |
| `ShowHelp` | `bool` | Display help text |
| `ShowStatusBar` | `bool` | Display status bar |
| `EnableFiltering` | `bool` | Allow filtering items (requires `ShowHelp`) |
| `Styles` | `Styles` | Override default appearance |

#### Default

```go
entries := []choose.Item{
    {Name: "pizza", Desc: "It's always pizza time!"},
    {Name: "kebab", Desc: "I feel turkish today, kebab!"},
    {Name: "carbonara", Desc: "Carbonara, NO cream, please!"},
}

result, err := choose.Run(&choose.Config{
    Title:    "What do you wanna eat tonight?",
    ErrorMsg: "Please, select your meal.",
}, entries)
```

<details>
<summary><b>Custom styles</b></summary>

```go
var (
    amber  = compat.AdaptiveColor{Light: lipgloss.Color("#f59e0b"), Dark: lipgloss.Color("#fbbf24")}
    purple = compat.AdaptiveColor{Light: lipgloss.Color("#7e22ce"), Dark: lipgloss.Color("#a855f7")}
    green  = compat.AdaptiveColor{Light: lipgloss.Color("#166534"), Dark: lipgloss.Color("#22c55e")}
)

entries := []choose.Item{
    {Name: "pizza", Desc: "It's always pizza time!"},
    {Name: "kebab", Desc: "I feel turkish today, kebab!"},
    {Name: "carbonara", Desc: "Carbonara, NO cream, please!"},
}

result, err := choose.Run(&choose.Config{
    Title:    "What do you wanna eat tonight?",
    ErrorMsg: "Please, select your meal.",
    ShowHelp: true,
    Styles: choose.Styles{
        PrefixIcon:        "★",
        TitleStyle:        lipgloss.NewStyle().Background(green).Foreground(purple).Padding(0, 1),
        TitleBarStyle:     lipgloss.NewStyle(),
        ItemIcon:          "#",
        ItemStyle:         lipgloss.NewStyle().Foreground(amber),
        SelectedItemStyle: lipgloss.NewStyle().Foreground(purple),
    },
}, entries)
```

</details>

---

### Confirm

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/confirm-default.gif" alt="Confirm example">
</details>

Yes/no confirmation prompt with two visual modes:

- **`ModeDialog`** (default) - a bordered dialog box with an optional message body
- **`ModeInline`** - a compact single-line toggle with cursor and divider

| Field | Type | Description |
|-------|------|-------------|
| `Mode` | `Mode` | `ModeDialog` (default) or `ModeInline` |
| `Message` | `string` | Body text inside the dialog border (dialog only) |
| `Question` | `string` | Prompt text shown to the user |
| `OkButtonLabel` | `string` | Affirmative option label (default "Yes") |
| `CancelButtonLabel` | `string` | Negative option label (default "No") |
| `Cursor` | `string` | Cursor character before toggle options (inline only, default ">") |
| `Divider` | `string` | Separator between toggle options (inline only, default "/") |
| `Styles` | `Styles` | Override default appearance |

#### Dialog mode (default)

```go
result, err := confirm.Run(&confirm.Config{Question: "Continue?"})
```

#### Inline mode

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/confirm-inline-default.gif" alt="Confirm inline example">
</details>

```go
result, err := confirm.Run(&confirm.Config{
    Mode:     confirm.ModeInline,
    Question: "Continue?",
})
```

<details>
<summary><b>Dialog mode - custom styles</b></summary>

```go
var (
    cyan  = compat.AdaptiveColor{Light: lipgloss.Color("#4f46e5"), Dark: lipgloss.Color("#c7d2fe")}
    green = compat.AdaptiveColor{Light: lipgloss.Color("#166534"), Dark: lipgloss.Color("#22c55e")}
)

infoText := `Lorem ipsum dolor sit amet,
consectetur adipiscing elit %s...`

Green   := lipgloss.NewStyle().Foreground(green).Render
message := fmt.Sprintf(infoText, Green("elit"))

result, err := confirm.Run(&confirm.Config{
    Message:  message,
    Question: "Continue?",
    Styles: confirm.Styles{
        Width:       60,
        BorderColor: cyan,
    },
})
```

</details>

<details>
<summary><b>Inline mode - custom styles</b></summary>

```go
var (
    cyan  = compat.AdaptiveColor{Light: lipgloss.Color("#4f46e5"), Dark: lipgloss.Color("#c7d2fe")}
    green = compat.AdaptiveColor{Light: lipgloss.Color("#166534"), Dark: lipgloss.Color("#22c55e")}
    red   = compat.AdaptiveColor{Light: lipgloss.Color("#ef4444"), Dark: lipgloss.Color("#ef4444")}
)

result, err := confirm.Run(&confirm.Config{
    Mode:              confirm.ModeInline,
    Question:          "How do you feel?",
    OkButtonLabel:     "I'm super ok",
    CancelButtonLabel: "Next question, please!",
    Divider:           "|",
    Styles: confirm.Styles{
        PrefixIcon:        "★",
        PrefixIconColor:   red,
        DialogStyle:       lipgloss.NewStyle().Margin(1, 0),
        ButtonStyle:       lipgloss.NewStyle().Bold(true).Foreground(cyan),
        ActiveButtonStyle: lipgloss.NewStyle().Foreground(green),
    },
})
```

</details>

---

### Detail

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/detail-default.gif" alt="Detail example">
</details>

Collapsible detail/summary section, similar to the HTML `<details>` element. Displays a summary line with an expand/collapse indicator and toggles content visibility on user interaction.

| Field | Type | Description |
|-------|------|-------------|
| `Summary` | `string` | The always-visible summary line |
| `Content` | `string` | Collapsible content (supports multiline) |
| `CollapsedIndicator` | `string` | Indicator when collapsed (default "▶") |
| `ExpandedIndicator` | `string` | Indicator when expanded (default "▼") |
| `Styles` | `Styles` | Override default appearance |

#### Default

```go
expanded, err := detail.Run(&detail.Config{
    Summary: "What is prompti?",
    Content: "prompti is a collection of interactive terminal UI prompts\nbuilt on the charmbracelet bubbletea framework.",
})
```

#### Custom indicators

```go
expanded, err := detail.Run(&detail.Config{
    Summary:            "Configuration Options",
    Content:            "verbose: true\nlog_level: debug\nmax_retries: 3",
    CollapsedIndicator: "[+]",
    ExpandedIndicator:  "[-]",
})
```

#### Multiline content

```go
expanded, err := detail.Run(&detail.Config{
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
})
```

<details>
<summary><b>Custom styles</b></summary>

```go
var (
    cyan   = compat.AdaptiveColor{Light: lipgloss.Color("#4f46e5"), Dark: lipgloss.Color("#c7d2fe")}
    green  = compat.AdaptiveColor{Light: lipgloss.Color("#166534"), Dark: lipgloss.Color("#22c55e")}
    red    = compat.AdaptiveColor{Light: lipgloss.Color("#ef4444"), Dark: lipgloss.Color("#ef4444")}
    purple = compat.AdaptiveColor{Light: lipgloss.Color("#7e22ce"), Dark: lipgloss.Color("#a855f7")}
)

expanded, err := detail.Run(&detail.Config{
    Summary: "Release Notes v0.3.0",
    Content: "- Added detail/summary component\n- Improved theme system\n- Bug fixes and performance improvements",
    Styles: detail.Styles{
        PrefixIcon:      "★",
        PrefixIconColor: red,
        SummaryStyle:    lipgloss.NewStyle().Bold(true).Foreground(purple).PaddingRight(1),
        IndicatorStyle:  lipgloss.NewStyle().Foreground(green).PaddingRight(1),
        ContentStyle:    lipgloss.NewStyle().Foreground(cyan).PaddingLeft(3).MarginTop(1),
        DialogStyle:     lipgloss.NewStyle().Margin(1, 0),
    },
})
```

</details>

---

### ProgressBar

<details>
<summary><b>Preview</b></summary>

<img src="https://raw.githubusercontent.com/indaco/gh-assets/main/prompti/progressbar-default.gif" alt="ProgressBar example">
</details>

Animated progress meter that iterates over a list of items. Standalone-only - does not support `huh.Field` integration.

| Field | Type | Description |
|-------|------|-------------|
| `Items` | `[]string` | List of items to process |
| `OnProgressCmd` | `func(string) tea.Cmd` | Custom command per item (default: simulated delay) |
| `OnProgressMsg` | `string` | Label shown during progress |
| `OnCompletedMsg` | `string` | Label shown on completion (default "Done!") |
| `RunConcurrently` | `bool` | Process items concurrently via `tea.Batch` |
| `Styles` | `Styles` | Override default appearance |

#### Default

```go
fruits := []string{"apple", "banana", "orange", "grapes", "strawberry", "mango", "kiwi", "pear", "cherry"}

err := progressbar.Run(&progressbar.Config{Items: fruits})
```

#### Concurrent execution

```go
fruits := []string{"apple", "banana", "orange", "grapes", "strawberry", "mango", "kiwi", "pear", "cherry"}

err := progressbar.Run(&progressbar.Config{
    Items:           fruits,
    RunConcurrently: true,
})
```

<details>
<summary><b>Custom styles</b></summary>

```go
fruits := []string{"apple", "banana", "orange", "grapes", "strawberry", "mango", "kiwi", "pear", "cherry"}

err := progressbar.Run(&progressbar.Config{
    Items:         fruits,
    OnProgressMsg: "Eating:",
    Styles: progressbar.Styles{
        CurrentItemStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("411")),
        ShowLabel:        true,
        GradientFrom:     lipgloss.Color("#FF7CCB"),
        GradientTo:       lipgloss.Color("#FDFF8C"),
    },
})
```

</details>

## Usage Modes

Every prompti component can be used in two ways:

### Standalone

Call `Run()` directly for single-prompt use cases. This starts its own bubbletea program and returns the result:

```go
result, err := input.Run(&input.Config{
    Message: "Project name?",
})
```

### As a huh.Field

Use `NewField()` to create a [`huh.Field`](https://github.com/charmbracelet/huh) adapter that plugs into multi-step forms:

```go
var name string
var confirmed bool

form := huh.NewForm(
    huh.NewGroup(
        input.NewField(&input.Config{
            Message: "Project name?",
        }, &name),
        confirm.NewField(&confirm.Config{
            Question: "Create it?",
        }, &confirmed),
    ),
)

err := form.Run()
```

Form components (`input`, `choose`, `confirm`, `detail`) support both modes. Feedback components like `progressbar` are standalone-only.

## Examples

Complete runnable examples for every component and mode are available in the [examples](examples/) directory:

| Directory | Description |
|-----------|-------------|
| [`01-input`](examples/01-input) | Text input with validation, password, custom styles |
| [`02-choose`](examples/02-choose) | Single-select list with custom styles |
| [`03-confirm`](examples/03-confirm) | Dialog and inline confirm modes |
| [`04-detail`](examples/04-detail) | Collapsible detail/summary sections |
| [`05-progressbar`](examples/05-progressbar) | Progress bar with sequential and concurrent modes |
| [`06-huh-form`](examples/06-huh-form) | Multi-group form using `huh.Field` adapters |

## Related

- [bubbletea](https://github.com/charmbracelet/bubbletea) - the underlying TUI framework
- [huh](https://github.com/charmbracelet/huh) - form library from Charm (prompti provides `huh.Field` adapters)
- [lipgloss](https://github.com/charmbracelet/lipgloss) - style engine used for all visual customization

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

---

Made with [Charm](https://charm.sh).

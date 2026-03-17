<h1 align="center">
    <span style="font-family: monospace;">prompti</span>
</h1>
<h2 align="center">
Lightweight and customizable interactive prompt components for Go based CLI.
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
  <a href="https://github.com/indaco/prompti/security" target="_blank">
    <img src="https://img.shields.io/badge/security-govulncheck-green" alt="Security Scan" />
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

- **5 prompt types**: input, choose, confirm, toggle, progressbar
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

`input` is a text input prompt supporting default values, validation (`type ValidateFunc func(string) error`), password echo mode, and customizable styles.

Built-in validation rules:

- alphanumeric
- digits only
- integers
- floats
- email address
- URL

#### Default

[Source Code](_examples/input/default/main.go)

<img src="https://statics.sveltin.io/github/prompti/input/input-default.gif" alt="Input example">

#### Default value and validation

[Source Code](_examples/input/initial-value/main.go)

<img src="https://statics.sveltin.io/github/prompti/input/input-initial-value.gif" alt="Input example with default value and validation">

#### Custom styles

[Source Code](_examples/input/custom-styles/main.go)

<img src="https://statics.sveltin.io/github/prompti/input/input-styled.gif" alt="Input example with custom styles">

#### Other Examples

- [Email](_examples/input/email/main.go)
- [Password](_examples/input/password/main.go)

### Choose

`choose` is a single-select list prompt for browsing a set of items.

#### Default

[Source Code](_examples/choose/default/main.go)

<img src="https://statics.sveltin.io/github/prompti/choose/choose-default.gif" alt="Choose example">

#### Custom styles

[Source Code](_examples/choose/custom-styles/main.go)

<img src="https://statics.sveltin.io/github/prompti/choose/choose-styled.gif" alt="Choose example with custom styles">

### Confirm

`confirm` is a yes/no confirmation dialog rendered in a styled box.

#### Default

[Source Code](_examples/confirm/default/main.go)

<img src="https://statics.sveltin.io/github/prompti/confirm/confirm-default.gif" alt="Confirm example">

#### Custom styles

[Source Code](_examples/confirm/custom-styles/main.go)

<img src="https://statics.sveltin.io/github/prompti/confirm/confirm-styled.gif" alt="Confirm example with custom styles">

### Toggle

`toggle` is an inline yes/no prompt. It works like `confirm` but renders the options inline rather than in a box.

#### Default

[Source Code](_examples/toggle/default/main.go)

<img src="https://statics.sveltin.io/github/prompti/toggle/toggle-default.gif" alt="Toggle example">

#### Custom styles

[Source Code](_examples/toggle/custom-styles/main.go)

<img src="https://statics.sveltin.io/github/prompti/toggle/toggle-styled.gif" alt="Toggle example with custom styles">

### ProgressBar

`progressbar` is an animated progress meter that iterates over a list of items.

#### Default

[Source Code](_examples/progressbar/default/main.go)

<img src="https://statics.sveltin.io/github/prompti/progressbar/progressbar-default.gif" alt="ProgressBar example">

#### Custom styles

[Source Code](_examples/progressbar/custom-styles/main.go)

<img src="https://statics.sveltin.io/github/prompti/progressbar/progressbar-styled.gif" alt="ProgressBar styled example">

#### Concurrent execution

[Source Code](_examples/progressbar/run-batch/main.go)

## License

prompti is free and open-source software licensed under the MIT License.

---

Made with [Charm](https://charm.sh).

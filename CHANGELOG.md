# Changelog

All notable changes to this project will be documented in this file.

This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html). The changelog is generated and managed by [sley](https://github.com/indaco/sley).

## 0.2.5 - 2023-09-28

### Revert

- `uniseg` to v0.2.0

## 0.2.4 - 2023-09-28

### Fixes

- **input**: remove unused `BackgroundStyle` in TextInput as per bubbles [v0.16.1](https://github.com/charmbracelet/bubbles/releases/tag/v0.16.1)
- **input**: replace deprecated `CursorStyle`in favour of Cursor.Style as per bubbles [v0.16.1](https://github.com/charmbracelet/bubbles/releases/tag/v0.16.1)

### Chores

- update [lipgloss](https://github.com/charmbracelet/lipgloss) to v0.8.0
- update [bubbles](github.com/charmbracelet/bubbles) to v0.16.1
- update [bubbletea](github.com/charmbracelet/bubbletea) to v0.24.2

## 0.2.3 - 2023-04-04

### Fixes

- **choose**: value of type func(s string) string as type func(strs ...string)"
- render TUI to stdout by default

### Chores

- update [lipgloss](https://github.com/charmbracelet/lipgloss) to v0.7.1

### CI

- update [golangci-lint](https://github.com/golangci/golangci-lint) to v1.52.2

# Variables

export LOG_STYLE := "emoji"
logger := "scripts/lib/logger.sh"

# Go commands

go := "go"
goclean := go + " clean"
vhs := "vhs"

# Default - show help
default:
    @just --list

# Clean the build directory and Go cache
clean:
    @. {{ logger }} && log_info "Clean the build directory and Go cache"
    rm -f coverage.out coverage.html
    {{ goclean }} -cache

# === Code Quality ===

# Format code
fmt:
    @. {{ logger }} && log_info "Running fmt and gofumpt"
    {{ go }} fmt ./...

# Run go-modernize with auto-fix
modernize:
    @. {{ logger }} && log_info "Running go-modernize"
    modernize --fix ./...

# Run golangci-lint
lint:
    @. {{ logger }} && log_info "Running golangci-lint"
    golangci-lint run ./...

# Run goreportcard-cli
reportcard:
    @. {{ logger }} && log_info "Running goreportcard-cli"
    goreportcard-cli -v

# Run govulncheck
security-scan:
    @. {{ logger }} && log_info "Running govulncheck"
    govulncheck ./...

# Run modernize, lint, and reportcard
check: fmt modernize lint reportcard

# Run go mod tidy
tidy:
    @. {{ logger }} && log_info "Running go mod tidy"
    {{ go }} mod tidy

# Run go mod download
deps:
    @. {{ logger }} && log_info "Running go mod download"
    {{ go }} mod download

# === Test Recipes ===

# Run all tests and print code coverage value
test:
    @. {{ logger }} && log_info "Run all tests"
    {{ go }} test $({{ go }} list ./... | grep -Ev 'internal/testutils|/examples/') -coverprofile=coverage.txt
    @. {{ logger }} && log_info "Total Coverage"
    {{ go }} tool cover -func=coverage.txt | grep total | awk '{print $3}'

# Clean go tests cache and run all tests
test-force:
    @. {{ logger }} && log_info "Clean go tests cache and run all tests"
    {{ go }} clean -testcache
    just test

# Run all tests and generate coverage report.
test-coverage:
    @. {{ logger }} && log_info "Run all tests and generate coverage report"
    {{ go }} test -count=1 -timeout 30s $({{ go }} list ./... | grep -Ev 'internal/testutils|/examples/') -covermode=atomic -coverprofile=coverage.txt

# Run all tests with race detector
test-race:
    @. {{ logger }} && log_info "Running tests with race detector"
    {{ go }} test -race $({{ go }} list ./... | grep -Ev 'internal/testutils|/examples/')

# === Utilities ===

# Update dependencies
deps-update:
    @. {{ logger }} && log_info "Running go update deps"
    {{ go }} get -u ./...
    {{ go }} mod tidy

# === Demos ===

# Record a single VHS tape
_record-tape name:
    {{ vhs }} examples/tapes/{{ name }}.tape

# Record all VHS tapes
demo-record:
    @. {{ logger }} && log_info "Recording all VHS tapes"
    for tape in examples/tapes/*.tape; do {{ vhs }} "$tape"; done
    @. {{ logger }} && log_success "All VHS tapes recorded"

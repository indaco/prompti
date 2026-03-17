#!/usr/bin/env bash
set -euo pipefail

# Generate demo GIFs for all examples using VHS.
# Usage: ./scripts/generate-gifs.sh [example_path]
#   No arguments: generate all GIFs
#   With argument: generate a single GIF (e.g. ./scripts/generate-gifs.sh input/default)

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
EXAMPLES_DIR="${PROJECT_ROOT}/_examples"

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BOLD='\033[1m'
RESET='\033[0m'

if ! command -v vhs &>/dev/null; then
    printf "${RED}error:${RESET} vhs is not installed. Install it from https://github.com/charmbracelet/vhs\n" >&2
    exit 1
fi

generate_gif() {
    local tape_dir="$1"
    local tape_file="${tape_dir}/demo.tape"
    local relative_path="${tape_dir#"${EXAMPLES_DIR}/"}"

    if [[ ! -f "${tape_file}" ]]; then
        printf "${YELLOW}skip:${RESET} %s (no demo.tape found)\n" "${relative_path}"
        return 0
    fi

    printf "${BOLD}generating:${RESET} %s ... " "${relative_path}"

    if (cd "${tape_dir}" && vhs demo.tape -q) 2>/dev/null; then
        printf "${GREEN}done${RESET}\n"
    else
        printf "${RED}failed${RESET}\n"
        return 1
    fi
}

main() {
    local failed=0

    if [[ $# -gt 0 ]]; then
        local target="${EXAMPLES_DIR}/$1"
        if [[ ! -d "${target}" ]]; then
            printf "${RED}error:${RESET} directory not found: %s\n" "${target}" >&2
            exit 1
        fi
        generate_gif "${target}" || failed=1
    else
        printf "${BOLD}Generating demo GIFs for all examples...${RESET}\n\n"

        while IFS= read -r tape_file; do
            generate_gif "$(dirname "${tape_file}")" || ((failed++))
        done < <(find "${EXAMPLES_DIR}" -name "demo.tape" | sort)

        printf "\n${BOLD}Done.${RESET}"
        if [[ ${failed} -gt 0 ]]; then
            printf " ${RED}%d failed.${RESET}\n" "${failed}"
            exit 1
        else
            printf " ${GREEN}All succeeded.${RESET}\n"
        fi
    fi
}

main "$@"

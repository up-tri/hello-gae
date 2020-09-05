
main() {
  go run ${PROJ_ROOT}/main.go
}

SCRIPT_DIR="$(cd "$(dirname "${0}")" && echo "${PWD}")"
PROJ_ROOT="${SCRIPT_DIR}/.."
main "$@"

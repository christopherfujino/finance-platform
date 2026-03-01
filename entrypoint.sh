#!/usr/bin/env bash
# Entrypoint for Docker container

set -euo pipefail

REPO_ROOT="$(dirname $(realpath ${BASH_SOURCE[0]}) )"

pushd "${REPO_ROOT}/go" >/dev/null

fallback_to_shell() {
  echo "Trapped SIGINT. Falling back to /bin/bash..." 1>&2
  exec /bin/bash
}

trap fallback_to_shell INT

if [[ -z "$@" ]]; then
  go run . ../fake.csv
else
  $@
fi

popd

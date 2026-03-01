#!/usr/bin/env bash

# Entrypoint for Docker container

REPO_ROOT="$(dirname $(realpath ${BASH_SOURCE[0]}) )"

pushd "${REPO_ROOT}/go" 2>/dev/null

if [[ -z "$@" ]]; then
  go run . ../fake.csv
else
  $@
fi

popd

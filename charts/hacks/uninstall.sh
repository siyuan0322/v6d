#!/bin/bash

ROOT=$(dirname "${BASH_SOURCE[0]}")/..

helm uninstall vineyard -n vineyard-system

kubectl -n vineyard-system delete localobjects --all || true
kubectl -n vineyard-system delete globalobjects --all || true

set +x
set +e
set +o pipefail

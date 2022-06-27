#!/usr/bin/env bash

# based on https://docs.docker.com/compose/startup-order/

set -e

readonly SVC_NAME=accountapi

waitForSvc() {
  until curl -X GET "http://${1}:8080/v1/health" --connect-timeout 1 &> /dev/null; do
    echo "${SVC_NAME} service is not yet available - sleeping"
    sleep 1
  done
}

waitForSvc "${SVC_NAME}"
echo "${SVC_NAME} service is up - executing tests"

#go test -tags e2e -coverpkg=./... ./... -covermode=atomic
go test \
  -tags e2e \
  -coverpkg ./... \
  -covermode=atomic \
  ./...
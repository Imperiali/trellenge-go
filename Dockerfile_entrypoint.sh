#!/bin/bash

echo "Starting development entrypoint..."

if [ ! -d vendor ]; then
  go mod vendor
fi

PORT=8888 go run cmd/worker/main.go &
modd
#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

mkdir -p out

GOOS=darwin  GOARCH=amd64 go build -o out/slack2md.darwin-amd64.out
GOOS=linux   GOARCH=amd64 go build -o out/slack2md.linux-amd64.out
GOOS=windows GOARCH=amd64 go build -o out/slack2md.windows-amd64.exe

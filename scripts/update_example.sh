#!/bin/bash

git checkout -B example
go build
./slack2md -input testdata -messages 4
git add -f output
git commit -m "example"

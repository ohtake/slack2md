#!/bin/bash

git checkout -B example
go build
./slack2md -input test_data -messages 4
git add -f output
git commit -m "example"

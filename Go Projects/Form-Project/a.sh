#!/bin/bash

export CGO_ENABLED=1
export GO111MODULE=on
export GOARCH=amd64
export GOOS=linux

go build -o myapp

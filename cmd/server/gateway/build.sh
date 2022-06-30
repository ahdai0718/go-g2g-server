#!/bin/sh

env GOOS=linux GOARCH=amd64 go build -o gatewayserveramd64
#env GOOS=windows GOARCH=amd64 go build -o gatewayserverwindowsamd64
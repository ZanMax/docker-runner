#!/bin/bash
# Build the project for different platforms

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o bin/docker-runner-linux docker-runner.go

# Build for MacOS
GOOS=darwin GOARCH=amd64 go build -o bin/docker-runner-macos docker-runner.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o bin/docker-runner-windows.exe docker-runner.go
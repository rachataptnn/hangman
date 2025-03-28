#!/bin/bash
GOOS=linux GOARCH=amd64 go build -o hangman-linux
GOOS=windows GOARCH=amd64 go build -o hangman-windows.exe
GOOS=darwin GOARCH=amd64 go build -o hangman-mac-int
GOOS=darwin GOARCH=arm64 go build -o hangman-mac-arm
echo "Build completed!"

#!/bin/bash

rm -rf build
GOOS=darwin GOARCH=amd64 go build -o build/swnt.osx
GOOS=linux GOARCH=amd64 go build -o build/swnt.linux
tar czvf build/swnt.tar.gz build/swnt.*

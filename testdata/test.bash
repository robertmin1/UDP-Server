#!/usr/bin/env bash

set -euo pipefail
shopt -s nullglob globstar

git clone https://github.com/robertmin1/UDP-Server && cd UDP-Server
cd testdata
go run make-file.go
cd ..

go run UDP-server.go

if grep -i "hello world" testdata/text.txt; then
    exit 0
else
    exit 1
fi

#!/usr/bin/env bash

set -euo pipefail
shopt -s nullglob globstar

git clone https://github.com/robertmin1/UDP-Sever && cd UDP-Sever
cd testdata
go run make-file.go
cd ..
./t.bash
timeout 10s go run UDP-sever.go

if grep -i "hello world" /testdata/test.txt; then
    return 0
else
    return 1
fi
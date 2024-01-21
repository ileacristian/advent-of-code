#!/usr/bin/env bash

for day in day*/ ; do
    echo "Running ${day}"
    cd "${day}"
    go run main.go
    cd .. # TODO: see if this can be improved
done
#!/usr/bin/env bash

SESSION_COOKIE="your-session-cookie"

for day in {01..25}; do
    mkdir -p "day${day}"

    URL="https://adventofcode.com/2024/day/${day#0}/input"  # the '#0' part removes the leading zero for the day

    wget --header "Cookie: session=${SESSION_COOKIE}" -O "day${day}/day${day}.txt" "$URL"
done
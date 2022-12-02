#!/bin/bash

FILENAME="day01.txt"

CALORIES_SUMS=()
CURRENT_SUM=0

while read LINE
do
    if [[ "${LINE}" != "" ]]
    then
        CURRENT_SUM=$(( ${CURRENT_SUM} + ${LINE} ))
    else
        CALORIES_SUMS+=(${CURRENT_SUM})
        CURRENT_SUM=0
    fi
done < "${FILENAME}"

IFS=$'\n'
    SORTED=$(sort -n <<<"${CALORIES_SUMS[*]}")
    MAX=$(echo "${SORTED}" | tail -1)
unset IFS

# part 1
echo "Elf carrying the most Calories. How many total Calories is that Elf carrying? ${MAX}"
echo

TOP3=($(echo "${SORTED}" | tail -3))

#part 2
echo "Top three Elves carrying the most Calories. How many Calories are those Elves carrying in total? $(( ${TOP3[0]} + ${TOP3[1]} + ${TOP3[2]} ))"

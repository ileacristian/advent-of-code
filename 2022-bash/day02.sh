#!/bin/bash

FILENAME="day02.txt"

ROCK="A"
PAPER="B"
SCISSORS="C"

TOTAL_SCORE=0

while read LINE || [[ -n "${LINE}" ]]
do
    FIRST_HAND=$(cut -d " " -f 1 <<< "${LINE}")
    SECOND_HAND=$(cut -d " " -f 2 <<< "${LINE}")

    # echo "${FIRST_HAND} ${SECOND_HAND}"

    SECOND_HAND=$(tr "X" "A" <<< "${SECOND_HAND}")
    SECOND_HAND=$(tr "Y" "B" <<< "${SECOND_HAND}")
    SECOND_HAND=$(tr "Z" "C" <<< "${SECOND_HAND}")

    # echo "${FIRST_HAND} ${SECOND_HAND}"

    SCORE=0

    if [[ "${SECOND_HAND}" = $ROCK ]]
    then
        SCORE=1
    elif [[ "${SECOND_HAND}" = $PAPER ]]
    then
        SCORE=2
    elif [[ "${SECOND_HAND}" = $SCISSORS ]]
    then
        SCORE=3
    fi

    if [[ "${FIRST_HAND}" = "${SECOND_HAND}" ]]
    then
        # echo "DRAW"
        SCORE=$((${SCORE} + 3))  # DRAW
    elif [[ "${FIRST_HAND}" = "${PAPER}" && "${SECOND_HAND}" = "${ROCK}" ]] ||
         [[ "${FIRST_HAND}" = "${ROCK}" && "${SECOND_HAND}" = "${SCISSORS}" ]] ||
         [[ "${FIRST_HAND}" = "${SCISSORS}" && "${SECOND_HAND}" = "${PAPER}" ]]
    then
        # echo "LOST"
         SCORE=$((${SCORE} + 0)) # LOST
    else
        # echo "WIN"
         SCORE=$((${SCORE} + 6)) #WIN
    fi

    # echo "Score for this rount: ${SCORE}"

    TOTAL_SCORE=$((${TOTAL_SCORE} + ${SCORE}))
done < "${FILENAME}"

echo "Total score following the strategy guide is (part1) ${TOTAL_SCORE}"

TOTAL_SCORE=0
WIN="Z"
DRAW="Y"
LOSE="X"

function handPairForOpponentAndState() {
    OPPONENT="${1}"
    GAME_STATE="${2}"

    # RETURN_VALUE=0

    if [[ "${GAME_STATE}" = "${WIN}" ]]
    then
        # echo "WIN"
        if [[ "${OPPONENT}" = "${ROCK}" ]]
        then
            # echo "PAPER"
            echo 2 # PAPER
        elif [[ "${OPPONENT}" = "${PAPER}" ]]
        then
            # echo "SCISSORS"
            echo 3 # SCISSORS
        elif [[ "${OPPONENT}" = "${SCISSORS}" ]]
        then
            # echo "ROCK"
            echo 1 # ROCK
        fi
    elif [[ "${GAME_STATE}" = "${DRAW}" ]]
    then
        # echo "draw"
        if [[ "${OPPONENT}" = "${ROCK}" ]]
        then
            # echo "ROCK"
            echo 1
        elif [[ "${OPPONENT}" = "${PAPER}" ]]
        then
            # echo "PAPER"
            echo 2
        elif [[ "${OPPONENT}" = "${SCISSORS}" ]]
        then
            # echo "SCISSORS"
            echo 3
        fi
    elif [[ "${GAME_STATE}" = "${LOSE}" ]]
    then
        # echo "lose"
        if [[ "${OPPONENT}" = "${ROCK}" ]]
        then
            # echo "SCISSORS"
            echo 3 # SCISSORS
        elif [[ "${OPPONENT}" = "${PAPER}" ]]
        then
            # echo "ROCK"
            echo 1 # ROCK
        elif [[ "${OPPONENT}" = "${SCISSORS}" ]]
        then
            # echo "PAPER"
            echo 2 # PAPER
        fi
    fi
    # echo "${RETURN_VALUE}"
}

while read LINE || [[ -n "${LINE}" ]]
do
    FIRST=$(cut -d " " -f 1 <<< "${LINE}")
    SECOND=$(cut -d " " -f 2 <<< "${LINE}")

    SCORE=0

    if [[ "${SECOND}" = $WIN ]]
    then
        SCORE=6
    elif [[ "${SECOND}" = $DRAW ]]
    then
        SCORE=3
    elif [[ "${SECOND}" = $LOSe ]]
    then
        SCORE=0
    fi

    SCORE=$((${SCORE} + $(handPairForOpponentAndState "${FIRST}" "${SECOND}")))
    TOTAL_SCORE=$((${TOTAL_SCORE} + ${SCORE}))
done < "${FILENAME}"

echo "The socre following the strategy guide is (part2) ${TOTAL_SCORE}"
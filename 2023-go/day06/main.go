package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Race struct {
	RecordDistance int
	Time           int
}

func (r Race) String() string {
	return fmt.Sprintf("[%dms %dmm]", r.Time, r.RecordDistance)
}

func main() {
	file, err := os.Open("day06.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var races []Race
	lineScanned := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)[1:]

		if races == nil {
			races = make([]Race, len(fields))
		}

		for i, field := range fields {
			if lineScanned == 0 {
				races[i].Time, _ = strconv.Atoi(field)
			} else {
				races[i].RecordDistance, _ = strconv.Atoi(field)
			}
		}

		lineScanned++
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("First Part: ", FirstPart(races))
	fmt.Println("Second Part: ", SecondPart(races))
}

func FirstPart(races []Race) int {
	possibilities := 1
	for _, race := range races {
		possibilities *= race.WaysOfWinning()
	}
	return possibilities
}

// we'll merge all the races into one
func SecondPart(races []Race) int {
	raceTimeStr := ""
	raceRecordStr := ""
	for _, race := range races {
		raceTimeStr += fmt.Sprint(race.Time)
		raceRecordStr += fmt.Sprint(race.RecordDistance)
	}

	raceRecord, _ := strconv.Atoi(raceRecordStr)
	raceTime, _ := strconv.Atoi(raceTimeStr)

	theRace := Race{RecordDistance: raceRecord, Time: raceTime}

	return theRace.WaysOfWinning()
}

func (r Race) WaysOfWinning() int {
	result := 0
	for holdTime := 1; holdTime <= r.Time; holdTime++ {
		if r.Outcome(holdTime) > r.RecordDistance {
			result++
		}
	}
	return result
}

func (r Race) Outcome(holdTime int) int {
	speed := holdTime
	remainingRaceTime := r.Time - holdTime
	return speed * remainingRaceTime
}

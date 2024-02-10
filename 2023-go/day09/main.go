package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("day09.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	readings := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		readings = append(readings, ParseLine(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("First Part: ", FirstPart(readings))
	fmt.Println("Second Part: ", SecondPart(readings))
}

func ParseLine(line string) []int {
	fields := strings.Fields(line)
	result := make([]int, len(fields))

	for i, v := range fields {
		result[i], _ = strconv.Atoi(v)
	}

	return result
}

func FirstPart(readings [][]int) int {
	result := 0
	for _, reading := range readings {
		result += nextElement(reading)
	}
	return result
}

func SecondPart(readings [][]int) int {
	result := 0
	for _, reading := range readings {
		result += prevElement(reading)
	}
	return result
}

func nextElement(slice []int) int {
	list := make([]int, len(slice))
	copy(list, slice)

	result := list[len(list)-1]

	notThereYet := true
	for notThereYet {
		notThereYet = false
		newList := make([]int, len(list)-1)
		for i := 0; i < len(list)-1; i++ {
			newList[i] = list[i+1] - list[i]
			if newList[i] != 0 {
				notThereYet = true
			}
		}

		result += newList[len(newList)-1]
		list = newList
	}

	return result
}

func prevElement(slice []int) int {
	list := make([]int, len(slice))
	copy(list, slice)

	results := []int{list[0]}

	notThereYet := true
	for notThereYet {
		notThereYet = false
		newList := make([]int, len(list)-1)
		for i := 0; i < len(list)-1; i++ {
			newList[i] = list[i+1] - list[i]
			if newList[i] != 0 {
				notThereYet = true
			}
		}

		results = append(results, newList[0])
		list = newList
	}

	result := 0
	for i := len(results) - 2; i >= 0; i-- {
		result = results[i] - result
	}

	return result
}

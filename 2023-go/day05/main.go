package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("day05.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	seeds := make([]int, 0)

	seed_to_soil := Mapping{}
	soil_to_fertiziler := Mapping{}
	fertilizer_to_water := Mapping{}
	water_to_light := Mapping{}
	light_to_temperature := Mapping{}
	temperature_to_humidity := Mapping{}
	humidity_to_location := Mapping{}

	mapsByName := map[string]*Mapping{
		"seed-to-soil":            &seed_to_soil,
		"soil-to-fertilizer":      &soil_to_fertiziler,
		"fertilizer-to-water":     &fertilizer_to_water,
		"water-to-light":          &water_to_light,
		"light-to-temperature":    &light_to_temperature,
		"temperature-to-humidity": &temperature_to_humidity,
		"humidity-to-location":    &humidity_to_location,
	}

	mapsOrdered := []*Mapping{
		&seed_to_soil,
		&soil_to_fertiziler,
		&fertilizer_to_water,
		&water_to_light,
		&light_to_temperature,
		&temperature_to_humidity,
		&humidity_to_location,
	}

	currentMap := &seed_to_soil

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "seeds") {
			parts := strings.Split(line, " ")
			for _, part := range parts[1:] {
				num, _ := strconv.Atoi(part)
				seeds = append(seeds, num)
			}
		} else if strings.HasSuffix(line, "map:") {
			fields := strings.Fields(line)
			mapName := fields[0]
			currentMap = mapsByName[mapName]
		} else if line != "" && unicode.IsDigit(rune(line[0])) {
			FillMap(line, currentMap)
		} else {
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("First Part: ", FirstPart(seeds, mapsOrdered))
	// fmt.Println("Second Part: ", SecondPart(cards))
}

func FirstPart(seeds []int, maps []*Mapping) int {
	minLocation := 990000000

	for _, seed := range seeds {
		destination := seed
		for _, mapping := range maps {
			destination = (*mapping).GetMapping(destination)
		}

		if destination < minLocation {
			minLocation = destination
		}
	}

	return minLocation
}

func FillMap(line string, mapping *Mapping) {
	parts := strings.Fields(line)
	destStart, _ := strconv.Atoi(parts[0])
	sourceStart, _ := strconv.Atoi(parts[1])
	length, _ := strconv.Atoi(parts[2])

	mappingInstr := MappingInstruction{destStart: destStart, sourceStart: sourceStart, length: length}
	mapping.AddInstruction(mappingInstr)
}

type Mapping struct {
	instructions []MappingInstruction
}

func (m *Mapping) AddInstruction(instruction MappingInstruction) {
	m.instructions = append(m.instructions, instruction)
}

func (m *Mapping) GetMapping(source int) int {

	for _, mappingInstr := range m.instructions {
		if destination, found := mappingInstr.mapping(source); found {
			return destination
		}
	}

	return source
}

type MappingInstruction struct {
	destStart   int
	sourceStart int
	length      int
}

func (m MappingInstruction) mapping(source int) (destination int, found bool) {
	if source >= m.sourceStart && source <= m.sourceStart+m.length-1 {
		offset := source - m.sourceStart
		destination = m.destStart + offset
		found = true
	} else {
		destination = -1
		found = false
	}
	return
}

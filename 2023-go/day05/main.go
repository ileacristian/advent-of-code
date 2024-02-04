package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	fmt.Println("Second Part: ", SecondPart(seeds, mapsOrdered))
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

func SecondPart(seeds []int, maps []*Mapping) int {
	seedRanges := []Range{}
	for pair := 0; pair < len(seeds)/2; pair++ {
		seedRange := Range{
			Start:  seeds[pair*2],
			Length: seeds[pair*2+1],
		}

		seedRanges = append(seedRanges, seedRange)
	}

	destinations := []Range(seedRanges)
	for _, mapping := range maps {
		newDestinations := []Range{}
		for _, source := range destinations {
			newDestinations = append(newDestinations, (*mapping).GetMappingRanges(source)...)
		}
		destinations = newDestinations
	}

	sort.Sort(ByStart(destinations))

	return destinations[0].Start
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

func (m *Mapping) GetMappingRanges(sourceRange Range) []Range {
	destinations := []Range{}
	overlaps := []Range{}

	for _, mappingInstr := range m.instructions {
		if destination, overlap, foundOverlap := mappingInstr.mappingRange(sourceRange); foundOverlap {
			destinations = append(destinations, destination)
			overlaps = append(overlaps, overlap)
		}
	}

	destinations = append(destinations, UnmappedRanges(sourceRange, overlaps)...)

	return destinations
}

func UnmappedRanges(sourceRange Range, overlaps []Range) []Range {
	ranges := []Range{}
	sort.Sort(ByStart(overlaps))

	if len(overlaps) == 0 {
		return []Range{sourceRange}
	}

	var startPivot int
	var currentOverlapIdx int
	if sourceRange.Start < overlaps[0].Start {
		startPivot = sourceRange.Start
	} else {
		for i, overlap := range overlaps {
			if overlap.Start+overlap.Length-1 >= sourceRange.Start {
				startPivot = overlap.Start + overlap.Length
				currentOverlapIdx = i + 1
				break
			}
		}
	}

	for currentOverlapIdx < len(overlaps) {
		if overlaps[currentOverlapIdx].Start-startPivot <= 0 {
			currentOverlapIdx++
			continue
		} else {
			if overlaps[currentOverlapIdx].Start >= sourceRange.Start+sourceRange.Length {
				ranges = append(ranges, sourceRange)
				startPivot = sourceRange.Start + sourceRange.Length
				break
			} else {
				ranges = append(ranges, Range{Start: startPivot, Length: overlaps[currentOverlapIdx].Start - startPivot})

			}
			pivotCandidate := overlaps[currentOverlapIdx].Start + overlaps[currentOverlapIdx].Length
			if pivotCandidate > sourceRange.Start+sourceRange.Length {
				startPivot = sourceRange.Start + sourceRange.Length
				break
			}
			startPivot = pivotCandidate
			currentOverlapIdx++
		}
	}

	if startPivot < sourceRange.Start+sourceRange.Length {
		ranges = append(ranges, Range{Start: startPivot, Length: sourceRange.Start + sourceRange.Length - startPivot})
	}

	return ranges
}

type MappingInstruction struct {
	destStart   int
	sourceStart int
	length      int
}

type Range struct {
	Start  int
	Length int
}

func NewRange(start, end int) Range {
	return Range{Start: start, Length: end - start + 1}
}

func (r *Range) End() int {
	return r.Start + r.Length - 1
}

func (r *Range) Overlaps(other Range) bool {
	var a, b Range

	if r.Start < other.Start {
		a = *r
		b = other
	} else {
		a = other
		b = *r
	}

	return b.Start <= a.End()
}

type ByStart []Range

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Less(i, j int) bool { return a[i].Start < a[j].Start }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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

func (m MappingInstruction) mappingRange(sourceRange Range) (destination Range, overlap Range, foundOverlap bool) {
	end1 := sourceRange.Start + sourceRange.Length - 1
	end2 := m.sourceStart + m.length - 1
	overlapStart := max(sourceRange.Start, m.sourceStart)
	overlapEnd := min(end1, end2)

	if overlapStart <= overlapEnd {
		overlapLength := overlapEnd - overlapStart + 1
		destinationOffset := overlapStart - m.sourceStart
		destination = Range{
			Start:  m.destStart + destinationOffset,
			Length: overlapLength,
		}

		overlap = Range{
			Start:  overlapStart,
			Length: overlapLength,
		}
		foundOverlap = true
	}
	return
}

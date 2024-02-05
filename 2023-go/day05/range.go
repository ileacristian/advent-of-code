package main

import (
	"fmt"
	"sort"
)

type Range struct {
	Start  int
	Length int
}

func NewRange(start, end int) Range {
	return Range{Start: start, Length: end - start + 1}
}

func (r Range) End() int {
	return r.Start + r.Length - 1
}

func (r Range) Mergeable(other Range) bool {
	if r == other {
		return true
	}

	var a, b Range

	if r.Start < other.Start {
		a = r
		b = other
	} else {
		a = other
		b = r
	}

	return b.Start <= a.End()+1
}

func (r Range) Merge(other Range) Range {
	if !r.Mergeable(other) {
		panic("Ranges don't overlap or are not next to each other")
	}
	var a, b Range

	if r.Start < other.Start {
		a = r
		b = other
	} else {
		a = other
		b = r
	}

	start := a.Start
	end := b.End()

	if end < a.End() {
		end = a.End()
	}

	return NewRange(start, end)
}

func (r Range) String() string {
	return fmt.Sprintf("[%d...%d]", r.Start, r.End())
}

type RangeSlice []Range

func (r RangeSlice) AnyMergeable() bool {
	for i := 0; i < len(r); i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i].Mergeable(r[j]) {
				return true
			}
		}
	}

	return false
}

func (r RangeSlice) MergeMergeable() RangeSlice {
	if !r.AnyMergeable() {
		return r
	}

	i := 0
	for r.AnyMergeable() {
		sort.Sort(ByStart(r))

		if i >= len(r)-1 {
			i = 0
		}

		if r[i].Mergeable(r[i+1]) {
			merge := r[i].Merge(r[i+1])
			r[i] = merge
			r = append(r[:i+1], r[i+2:]...)
		} else {
			i++
		}
	}

	return r
}

type ByStart []Range

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Less(i, j int) bool { return a[i].Start < a[j].Start }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

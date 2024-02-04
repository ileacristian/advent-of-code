package main

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

func (r Range) Overlaps(other Range) bool {
	var a, b Range

	if r.Start < other.Start {
		a = r
		b = other
	} else {
		a = other
		b = r
	}

	return b.Start <= a.End()
}

type ByStart []Range

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Less(i, j int) bool { return a[i].Start < a[j].Start }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

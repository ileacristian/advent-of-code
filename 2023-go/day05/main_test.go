package main

import (
	"reflect"
	"testing"
)

func TestMappingRange(t *testing.T) {
	t.Run("Testing mappingRange with input from test file", func(t *testing.T) {
		m := MappingInstruction{DestStart: 98, SourceStart: 50, Length: 2}
		s := Range{Start: 50, Length: 2}
		gotDest, gotOverlap, gotFound := m.mappingRange(s)
		wantDest, wantOverlap, wantFound := Range{Start: 98, Length: 2}, Range{Start: 50, Length: 2}, true

		if gotDest != wantDest || gotOverlap != wantOverlap || gotFound != wantFound {
			t.Errorf("Expected %v %v %t but got %v %v %t", wantDest, wantOverlap, wantFound, gotDest, gotOverlap, gotFound)
		}
	})
	t.Run("Testing mappingRange with no overlap", func(t *testing.T) {
		m := MappingInstruction{DestStart: 98, SourceStart: 50, Length: 2}
		s := Range{Start: 52, Length: 2}
		gotDest, gotOverlap, gotFound := m.mappingRange(s)
		wantDest, wantOverlap, wantFound := Range{Start: 0, Length: 0}, Range{Start: 0, Length: 0}, false

		if gotDest != wantDest || gotOverlap != wantOverlap || gotFound != wantFound {
			t.Errorf("Expected %v %v %t but got %v %v %t", wantDest, wantOverlap, wantFound, gotDest, gotOverlap, gotFound)
		}
	})

	t.Run("Testing mappingRange with some overlap", func(t *testing.T) {
		m := MappingInstruction{DestStart: 500, SourceStart: 1, Length: 100}
		s := Range{Start: 50, Length: 100}
		gotDest, gotOverlap, gotFound := m.mappingRange(s)
		wantDest, wantOverlap, wantFound := Range{Start: 549, Length: 51}, Range{Start: 50, Length: 51}, true

		if gotDest != wantDest || gotOverlap != wantOverlap || gotFound != wantFound {
			t.Errorf("Expected %v %v %t but got %v %v %t", wantDest, wantOverlap, wantFound, gotDest, gotOverlap, gotFound)
		}
	})

	t.Run("Testing mappingRange with full overlap inside", func(t *testing.T) {
		m := MappingInstruction{DestStart: 111, SourceStart: 1, Length: 100}
		s := Range{Start: 50, Length: 2}
		gotDest, gotOverlap, gotFound := m.mappingRange(s)
		wantDest, wantOverlap, wantFound := Range{Start: 160, Length: 2}, Range{Start: 50, Length: 2}, true

		if gotDest != wantDest || gotOverlap != wantOverlap || gotFound != wantFound {
			t.Errorf("Expected %v %v %t but got %v %v %t", wantDest, wantOverlap, wantFound, gotDest, gotOverlap, gotFound)
		}
	})
}

func TestGetMappingRanges(t *testing.T) {
	t.Run("Testing Mapping with one MappingInstruction and full overlapping input Range", func(t *testing.T) {
		m := Mapping{}
		m.AddInstruction(MappingInstruction{DestStart: 98, SourceStart: 50, Length: 2})
		s := Range{Start: 50, Length: 2}

		got := m.GetMappingRanges(s)
		want := []Range{{Start: 98, Length: 2}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("Testing Mapping with one MappingInstruction and overlapping input Range", func(t *testing.T) {
		m := Mapping{}
		m.AddInstruction(MappingInstruction{DestStart: 98, SourceStart: 50, Length: 2})
		s := Range{Start: 50, Length: 5}

		got := m.GetMappingRanges(s)
		want := []Range{{Start: 98, Length: 2}, {Start: 52, Length: 3}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("Testing Mapping with multiple MappingInstruction and overlapping input Range", func(t *testing.T) {
		m := Mapping{}
		m.AddInstruction(MappingInstruction{DestStart: 98, SourceStart: 50, Length: 2})
		m.AddInstruction(MappingInstruction{DestStart: 50, SourceStart: 52, Length: 48})

		s := Range{Start: 79, Length: 14}

		got := m.GetMappingRanges(s)
		want := []Range{{Start: 77, Length: 14}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("Testing Mapping with multiple MappingInstruction and overlapping input Range", func(t *testing.T) {
		m := Mapping{}
		m.AddInstruction(MappingInstruction{DestStart: 15, SourceStart: 0, Length: 37})
		m.AddInstruction(MappingInstruction{DestStart: 52, SourceStart: 37, Length: 2})
		m.AddInstruction(MappingInstruction{DestStart: 0, SourceStart: 39, Length: 15})

		s := Range{Start: 77, Length: 14}

		got := m.GetMappingRanges(s)
		want := []Range{{Start: 77, Length: 14}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("Testing Mapping with multiple MappingInstruction and overlapping input Range", func(t *testing.T) {
		m := Mapping{}
		m.AddInstruction(MappingInstruction{DestStart: 15, SourceStart: 0, Length: 37})
		m.AddInstruction(MappingInstruction{DestStart: 52, SourceStart: 37, Length: 2})
		m.AddInstruction(MappingInstruction{DestStart: 0, SourceStart: 39, Length: 39})

		s := Range{Start: 77, Length: 14}

		got := m.GetMappingRanges(s)
		want := []Range{{Start: 38, Length: 1}, {Start: 78, Length: 13}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})
}

func TestUnmappedRanges(t *testing.T) {
	t.Run("Testing UnmappedRanges with source and one overlap at the end (already mapped range) - should return unmapped range", func(t *testing.T) {
		got := UnmappedRanges(Range{Start: 1, Length: 100}, []Range{{Start: 51, Length: 100}})
		want := []Range{{Start: 1, Length: 50}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("Testing UnmappedRanges with source and one overlap at the beginning (already mapped range) - should return unmapped range", func(t *testing.T) {
		got := UnmappedRanges(Range{Start: 100, Length: 100}, []Range{{Start: 51, Length: 100}})
		want := []Range{{Start: 151, Length: 49}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("Testing UnmappedRanges with source and multiple overlaps (already mapped ranges) - should return unmapped ranges", func(t *testing.T) {
		got := UnmappedRanges(Range{Start: 1, Length: 100}, []Range{{Start: 69, Length: 3}, {Start: 42, Length: 7}})
		want := []Range{{Start: 1, Length: 41}, {Start: 49, Length: 20}, {Start: 72, Length: 29}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("Testing UnmappedRanges with source and one overlap that already covers the whole source - should return empty slice", func(t *testing.T) {
		got := UnmappedRanges(Range{Start: 1, Length: 100}, []Range{{Start: 0, Length: 400}})
		want := []Range{}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("Testing UnmappedRanges with source and one overlap that doesn't cover anything - should return the source slice", func(t *testing.T) {
		got := UnmappedRanges(Range{Start: 1, Length: 100}, []Range{{Start: 200, Length: 50}})
		want := []Range{{Start: 1, Length: 100}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("Testing UnmappedRanges with source and one overlap - should return the remaining range", func(t *testing.T) {
		got := UnmappedRanges(Range{Start: 50, Length: 5}, []Range{{Start: 50, Length: 2}})
		want := []Range{{Start: 52, Length: 3}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})

	t.Run("Testing UnmappedRanges with source and no items int he overlap array - should return the source range", func(t *testing.T) {
		got := UnmappedRanges(Range{Start: 50, Length: 5}, []Range{})
		want := []Range{{Start: 50, Length: 5}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v but got %v", want, got)
		}
	})
}

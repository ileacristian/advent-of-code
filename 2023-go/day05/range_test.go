package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewRange(t *testing.T) {
	t.Run("Tests Range constructor", func(t *testing.T) {
		r := NewRange(10, 20)
		if r.Start != 10 || r.End() != 20 || r.Length != 11 {
			t.Errorf("Expected start=%d end=%d length=%d got start=%d end=%d length=%d", 10, 20, 11, r.Start, r.End(), r.Length)
		}
	})
}

func TestEnd(t *testing.T) {
	t.Run("Tests Range End method", func(t *testing.T) {
		r := NewRange(10, 20)
		if r.Start != 10 || r.End() != 20 || r.Length != 11 {
			t.Errorf("Expected start=%d end=%d length=%d got start=%d end=%d length=%d", 10, 20, 11, r.Start, r.End(), r.Length)
		}
	})
}

func TestOverlaps(t *testing.T) {
	t.Run("Tests Range Mergeable method: no overlap", func(t *testing.T) {
		a := NewRange(10, 20)
		b := NewRange(30, 50)
		if a.Mergeable(b) || b.Mergeable(a) {
			t.Errorf("Expected no overlap")
		}
	})

	t.Run("Tests Range Mergeable method: partial overlap left", func(t *testing.T) {
		a := NewRange(10, 20)
		b := NewRange(5, 15)
		if !a.Mergeable(b) || !b.Mergeable(a) {
			t.Errorf("Expected overlap")
		}
	})

	t.Run("Tests Range Mergeable method: partial overlap right", func(t *testing.T) {
		a := NewRange(10, 20)
		b := NewRange(15, 25)
		if !a.Mergeable(b) || !b.Mergeable(a) {
			t.Errorf("Expected overlap")
		}
	})

	t.Run("Tests Range Mergeable method: full overlap", func(t *testing.T) {
		a := NewRange(10, 10)
		b := NewRange(10, 10)
		if !a.Mergeable(b) || !b.Mergeable(a) {
			t.Errorf("Expected no overlap")
		}
	})

	t.Run("Tests Range Merge method - with overlap", func(t *testing.T) {
		a := NewRange(10, 20)
		b := NewRange(15, 25)
		c := NewRange(20, 30)

		if a.Merge(b) != NewRange(10, 25) || b.Merge(c) != NewRange(15, 30) || a.Merge(c) != NewRange(10, 30) {
			t.Errorf("Merging failed")
		}
	})

	t.Run("Tests Range Merge method - with overlap inside", func(t *testing.T) {
		a := NewRange(10, 20)
		b := NewRange(12, 15)

		expect := NewRange(10, 20)
		got := a.Merge(b)

		if got != expect {
			t.Errorf("Merging failed - got %v but expected %v", got, expect)
		}
	})

	t.Run("Tests Range Merge method - with no overlap should panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		a := NewRange(10, 20)
		b := NewRange(30, 45)

		_ = a.Merge(b)
	})

	t.Run("Tests Range String method", func(t *testing.T) {
		a := NewRange(10, 20)
		b := NewRange(15, 30)

		got := fmt.Sprint(a.Merge(b))
		expect := "[10...30]"

		if expect != got {
			t.Errorf("Expected %s but got %s", expect, got)
		}
	})

	t.Run("Tests RangeSlice AnyMergeable - method - overlaps", func(t *testing.T) {
		rangeSlice := RangeSlice{NewRange(10, 20), NewRange(99, 100), NewRange(15, 25)}

		expect := true
		got := rangeSlice.AnyMergeable()

		if expect != got {
			t.Errorf("Expected %t but got %t", expect, got)
		}
	})

	t.Run("Tests RangeSlice AnyMergeable - method - no overlaps", func(t *testing.T) {
		rangeSlice := RangeSlice{NewRange(10, 20), NewRange(99, 100), NewRange(1, 8)}

		expect := false
		got := rangeSlice.AnyMergeable()

		if expect != got {
			t.Errorf("Expected %t but got %t", expect, got)
		}
	})

	t.Run("Tests RangeSlice AnyMergeable - method - next to eachother", func(t *testing.T) {
		rangeSlice := RangeSlice{NewRange(10, 98), NewRange(99, 100)}

		expect := true
		got := rangeSlice.AnyMergeable()

		if expect != got {
			t.Errorf("Expected %t but got %t", expect, got)
		}
	})

	t.Run("Test RangeSlice MergeMergeable 1", func(t *testing.T) {
		rangeSlice := RangeSlice{NewRange(10, 98), NewRange(99, 100)}

		expect := RangeSlice{NewRange(10, 100)}
		got := rangeSlice.MergeMergeable()

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("Expected %v but got %v", expect, got)
		}
	})

	t.Run("Test RangeSlice MergeMergeable 2 ", func(t *testing.T) {
		rangeSlice := RangeSlice{NewRange(10, 98), NewRange(4, 5), NewRange(2, 8), NewRange(99, 100), NewRange(15, 20), NewRange(15, 20)}

		expect := RangeSlice{NewRange(2, 8), NewRange(10, 100)}
		got := rangeSlice.MergeMergeable()

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("Expected %v but got %v", expect, got)
		}
	})

	t.Run("Test RangeSlice MergeMergeable 3", func(t *testing.T) {
		rangeSlice := RangeSlice{NewRange(10, 12), NewRange(99, 100)}

		expect := RangeSlice{NewRange(10, 12), NewRange(99, 100)}
		got := rangeSlice.MergeMergeable()

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("Expected %v but got %v", expect, got)
		}
	})

	t.Run("Test RangeSlice MergeMergeable 4", func(t *testing.T) {
		rangeSlice := RangeSlice{NewRange(10, 12), NewRange(5, 100), NewRange(10, 20)}

		expect := RangeSlice{NewRange(5, 100)}
		got := rangeSlice.MergeMergeable()

		if !reflect.DeepEqual(expect, got) {
			t.Errorf("Expected %v but got %v", expect, got)
		}
	})
}

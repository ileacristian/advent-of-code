package main

import "testing"

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
	t.Run("Tests Range Overlaps method: no overlap", func(t *testing.T) {
		a := NewRange(10, 20)
		b := NewRange(30, 50)
		if a.Overlaps(b) || b.Overlaps(a) {
			t.Errorf("Expected no overlap")
		}
	})

	t.Run("Tests Range Overlaps method: partial overlap left", func(t *testing.T) {
		a := NewRange(10, 20)
		b := NewRange(5, 15)
		if !a.Overlaps(b) || !b.Overlaps(a) {
			t.Errorf("Expected overlap")
		}
	})

	t.Run("Tests Range Overlaps method: partial overlap right", func(t *testing.T) {
		a := NewRange(10, 20)
		b := NewRange(15, 25)
		if !a.Overlaps(b) || !b.Overlaps(a) {
			t.Errorf("Expected overlap")
		}
	})

	t.Run("Tests Range Overlaps method: full overlap", func(t *testing.T) {
		a := NewRange(10, 10)
		b := NewRange(10, 10)
		if !a.Overlaps(b) || !b.Overlaps(a) {
			t.Errorf("Expected no overlap")
		}
	})
}

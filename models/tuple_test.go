package models

import (
	"testing"
)

// test creation of a point
func TestPoint(t *testing.T) {
	p := Point(4, -4, 3)
	if p.x != 4 || p.y != -4 || p.z != 3 || p.w != 1 {
		t.Errorf("Point(4, -4, 3) = %v, want %v", p, Tuple{4, -4, 3, 1})
	}
}

// test creation of a vector
func TestVector(t *testing.T) {
	v := Vector(4, -4, 3)
	if v.x != 4 || v.y != -4 || v.z != 3 || v.w != 0 {
		t.Errorf("Vector(4, -4, 3) = %v, want %v", v, Tuple{4, -4, 3, 0})
	}
}

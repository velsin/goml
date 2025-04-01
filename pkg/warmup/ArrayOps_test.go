package warmup

import (
	"slices"
	"testing"
)

func TestGetIntersection(t *testing.T) {
	a1, b1, e1 := []int{1, 3, 5}, []int{2, 3, 4, 5}, []int{3, 5}
	a2, b2, e2 := []float64{1.23, 1.90, 2.3}, []float64{1.22, 1.89, 2.1}, []float64{}
	a3, b3, e3 := []rune{'k', 'r', 'ä'}, []rune{'ä', 'å'}, []rune{'ä'}

	for i, v := range GetIntersection(a1, b1) {
		if v != e1[i] {
			t.Errorf("Failed on %v, got element %v, expected %v", e1, v, e1[i])
		}
	}
	for i, v := range GetIntersection(a2, b2) {
		if v != e2[i] {
			t.Errorf("Failed on %v, got element %v, expected %v", e2, v, e2[i])
		}
	}
	for i, v := range GetIntersection(a3, b3) {
		if v != e3[i] {
			t.Errorf("Failed on %v, got element %v, expected %v", e3, v, e3[i])
		}
	}
}

func TestGetUnion(t *testing.T) {
	a1, b1, e1 := []int{1, 3, 5}, []int{2, 3, 4, 5}, []int{1, 3, 5, 2, 4}
	a2, b2, e2 := []float64{1.23, 1.90, 2.3}, []float64{1.22, 1.89, 2.1}, []float64{1.23, 1.90, 2.3, 1.22, 1.89, 2.1}
	a3, b3, e3 := []rune{'k', 'r', 'ä'}, []rune{'ä', 'å'}, []rune{'k', 'r', 'ä', 'å'}

	slices.Sort(e1)
	slices.Sort(e2)
	slices.Sort(e3)

	r1 := GetUnion(a1, b1)
	slices.Sort(r1)
	r2 := GetUnion(a2, b2)
	slices.Sort(r2)
	r3 := GetUnion(a3, b3)
	slices.Sort(r3)

	for i, v := range r1 {
		if v != e1[i] {
			t.Errorf("Failed on %v, got element %v, expected %v", e1, v, e1[i])
		}
	}
	for i, v := range r2 {
		if v != e2[i] {
			t.Errorf("Failed on %v, got element %v, expected %v", e2, v, e2[i])
		}
	}
	for i, v := range r3 {
		if v != e3[i] {
			t.Errorf("Failed on %v, got element %v, expected %v", e3, v, e3[i])
		}
	}

}

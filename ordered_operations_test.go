package set

import "testing"

func TestOrderedDiff(t *testing.T) {
	var (
		expect   = []int{1, 3, 5}
		unexpect = []int{2, 4, 6}
		a        = NewOrdered[int](1, 2, 3, 4, 5)
		b        = NewOrdered[int](2, 4, 6)
		diff     = ODiff(a, b)
	)

	for _, n := range expect {
		testOrdHas(t, diff, n, true)
	}
	for _, n := range unexpect {
		testOrdHas(t, diff, n, false)
	}
}

func TestOrderedUnion(t *testing.T) {
	var (
		expect = []int{1, 2, 3, 4, 5, 6}
		a      = NewOrdered[int](1, 2, 5, 6)
		b      = NewOrdered[int](2, 3, 4, 5)
		union  = OUnion(a, b)
	)

	for _, n := range expect {
		testOrdHas(t, union, n, true)
	}
}

func TestOrderedIntersection(t *testing.T) {
	var (
		expect    = []int{3, 4, 5}
		unexpect  = []int{-2, -1, 1, 2, 6, 7, 8, 9}
		a         = NewOrdered[int](1, 2, 3, 4, 5, 8, 9)
		b         = NewOrdered[int](-1, -2, 3, 4, 5, 6, 7)
		intersect = OIntersection(a, b)
	)

	for _, n := range expect {
		testOrdHas(t, intersect, n, true)
	}
	for _, n := range unexpect {
		testOrdHas(t, intersect, n, false)
	}
}

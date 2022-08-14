package set

import (
	"testing"
)

type intu int

func (a intu) Less(b intu) bool {
	return a < b
}

func TestSortedDiff(t *testing.T) {
	var (
		a    = New[intu](1, 2, 3, 6, 7, 8, 10, 11, 12)
		b    = New[intu](2, 3, 4, 5, 8, 9)
		diff = Diff(a, b)
	)

	testHas(t, diff, 1, true)
	testHas(t, diff, 6, true)
	testHas(t, diff, 7, true)
	testHas(t, diff, 10, true)
	testHas(t, diff, 11, true)
	testHas(t, diff, 12, true)
	testHas(t, diff, 2, false)
	testHas(t, diff, 3, false)
	testHas(t, diff, 4, false)
	testHas(t, diff, 5, false)
	testHas(t, diff, 8, false)
	testHas(t, diff, 9, false)
}

func TestSortedUnion(t *testing.T) {
	var (
		a     = New[intu](1, 2, 3, 7, 8, 9)
		b     = New[intu](3, 4, 5, 6, 7, 10, 11, 12)
		union = Union(a, b)
	)

	for i := 1; i <= 12; i++ {
		testHas(t, union, intu(i), true)
	}
}

func TestSortedIntersection(t *testing.T) {
	var (
		a         = New[intu](1, 2, 3, 4, 5)
		b         = New[intu](3, 4, 5, 6, 7)
		intersect = Intersection(a, b)
	)

	testHas(t, intersect, 3, true)
	testHas(t, intersect, 4, true)
	testHas(t, intersect, 5, true)
}
